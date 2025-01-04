package middleware

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/y-nosuke/sample-task-api-go/app/framework/auth"
	fcontext "github.com/y-nosuke/sample-task-api-go/app/framework/context"
	ferrors "github.com/y-nosuke/sample-task-api-go/app/framework/errors"
	"golang.org/x/xerrors"
)

var keySet jwk.Set

func init() {
	fmt.Println("init auth.")

	jwksUrl := os.Getenv("AUTH_JWKS_URL")
	var err error
	if keySet, err = jwk.Fetch(context.Background(), jwksUrl); err != nil {
		panic(err)
	}
}

func ValidateTokenMiddleware() func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ectx echo.Context) error {
			cctx := fcontext.CastContext(ectx)
			fmt.Println("ValidateTokenMiddleware start. トークンを検証します。")

			tokenString := getToken(ectx.Request())
			if tokenString == "" {
				return ferrors.NewUnauthorizedError("認証されていません。 missing Authorization Header")
			}

			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
				kid, ok := token.Header["kid"].(string)
				if !ok {
					return nil, xerrors.Errorf("kid not found in token header")
				}

				publicKey, err := getPublicKey(kid)
				if err != nil {
					return nil, xerrors.Errorf("unable to get the public key. Error: %s", err.Error())
				}

				if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
					return nil, xerrors.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return publicKey, nil
			})
			if err != nil {
				fmt.Println(err.Error())
				return ferrors.NewUnauthorizedError("認証されていません。 invalid token")
			}

			auths, err := auth.NewAuth(token)
			if err != nil {
				fmt.Println(err.Error())
				return ferrors.NewUnauthorizedError("認証されていません。")
			}

			auth.SetAuth(cctx, auths)

			if err = next(cctx); err != nil {
				return xerrors.Errorf("next(): %w", err)
			}

			return nil
		}
	}

}

func getToken(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return ""
	}

	tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
	if tokenString == "" {
		return ""
	}

	return tokenString
}

func getPublicKey(kid string) (interface{}, error) {
	key, ok := keySet.LookupKeyID(kid)
	if !ok {
		return nil, fmt.Errorf("key not found in key set")
	}

	var publicKey interface{}
	if err := key.Raw(&publicKey); err != nil {
		return nil, err
	}

	return publicKey, nil
}
