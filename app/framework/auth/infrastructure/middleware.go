package infrastructure

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/lestrrat-go/jwx/jwk"
	fauth "github.com/y-nosuke/sample-task-api-go/app/framework/auth"
	fcontext "github.com/y-nosuke/sample-task-api-go/app/framework/context/infrastructure"
	ferrors "github.com/y-nosuke/sample-task-api-go/app/framework/errors"
	"golang.org/x/xerrors"
	"net/http"
	"os"
	"strings"
)

type ctxKey int

const (
	AUTH ctxKey = iota
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

func ValidateTokenMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ectx echo.Context) error {
		fmt.Println("トークン検証を実行します。")
		tokenString := getToken(ectx.Request())
		if tokenString == "" {
			return ferrors.New(ferrors.Unauthorized, "認証されていません。", fmt.Errorf("missing Authorization Header"))
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			kid, ok := token.Header["kid"].(string)
			if !ok {
				return nil, xerrors.Errorf("kid not found in token header")
			}

			publicKey, err := getPublicKey(kid)
			if err != nil {
				return nil, fmt.Errorf("unable to get the public key. Error: %s", err.Error())
			}

			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return publicKey, nil
		})
		if err != nil {
			return ferrors.New(ferrors.Unauthorized, "認証されていません。", fmt.Errorf("invalid token"))
		}

		auth, err := fauth.NewAuth(token)
		if err != nil {
			return ferrors.New(ferrors.Unauthorized, "認証されていません。", err)
		}

		cctx := fcontext.Cctx(ectx)
		cctx.WithValue(AUTH, auth)

		return next(ectx)
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
