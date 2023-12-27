package infrastructure

import (
	"context"
	"fmt"
	"github.com/y-nosuke/sample-task-api-go/app/framework/errors"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/y-nosuke/sample-task-api-go/app/framework/auth"
	"github.com/y-nosuke/sample-task-api-go/app/framework/auth/application/presenter"
	fcontext "github.com/y-nosuke/sample-task-api-go/app/framework/context"
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

func ValidateTokenMiddleware(authHandlerPresenter presenter.AuthHandlerPresenter) func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ectx echo.Context) error {
			fmt.Println("トークン検証を実行します。")
			cctx := fcontext.Cctx(ectx)

			tokenString := getToken(ectx.Request())
			if tokenString == "" {
				return authHandlerPresenter.Unauthorized(cctx.Ctx, "認証されていません。 missing Authorization Header")
			}

			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
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
				if err := authHandlerPresenter.Unauthorized(cctx.Ctx, "認証されていません。 invalid token"); err != nil {
					fmt.Println(1)
					return errors.SystemErrorf("taskPresenter.Forbidden(): %w", err)
				}
				fmt.Println(2)
				return errors.BusinessErrorf("taskPresenter.Forbidden()")
			}

			auths, err := auth.NewAuth(token)
			if err != nil {
				if err := authHandlerPresenter.Unauthorized(cctx.Ctx, "認証されていません。"); err != nil {
					return errors.SystemErrorf("taskPresenter.Forbidden(): %w", err)
				}
				return errors.BusinessErrorf("taskPresenter.Forbidden()")
			}

			auth.SetAuth(cctx, auths)

			return next(ectx)
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
		return nil, xerrors.Errorf("key not found in key set")
	}

	var publicKey interface{}
	if err := key.Raw(&publicKey); err != nil {
		return nil, xerrors.Errorf("key.Raw()")
	}

	return publicKey, nil
}
