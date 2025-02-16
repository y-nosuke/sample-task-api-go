package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/y-nosuke/sample-task-api-go/app/framework/auth"
	fcontext "github.com/y-nosuke/sample-task-api-go/app/framework/context"
	ferrors "github.com/y-nosuke/sample-task-api-go/app/framework/errors"
	"golang.org/x/xerrors"
)

func ValidateTokenMiddleware() func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ectx echo.Context) error {
			cctx := fcontext.CastContext(ectx)
			fmt.Println("ValidateTokenMiddleware start. トークンを検証します。")

			tokenString, err := getToken(ectx.Request())
			if err != nil {
				return ferrors.NewUnauthorizedErrorf(err, "認証されていません。 missing Authorization Header")
			}

			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
				if token.Method != jwt.SigningMethodRS256 {
					return nil, xerrors.Errorf("unexpected signing method: %v", token.Header["alg"])
				}

				kid, ok := token.Header["kid"].(string)
				if !ok || kid == "" {
					return nil, xerrors.Errorf("invalid or missing 'kid' in token header")
				}

				publicKey, err := getPublicKey(ectx.Request().Context(), kid)
				if err != nil {
					return nil, xerrors.Errorf("unable to get the public key. Error: %w", err)
				}

				return publicKey, nil
			})
			if err != nil {
				return ferrors.NewUnauthorizedErrorf(err, "認証されていません。 invalid token")
			}

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				if exp, ok := claims["exp"].(float64); ok {
					if time.Now().After(time.Unix(int64(exp), 0)) {
						return ferrors.NewUnauthorizedErrorf(err, "認証されていません。 token has expired")
					}
				}
				if nbf, ok := claims["nbf"].(float64); ok {
					if time.Now().Before(time.Unix(int64(nbf), 0)) {
						return ferrors.NewUnauthorizedErrorf(err, "認証されていません。 token is not yet valid")
					}
				}
			} else {
				return ferrors.NewUnauthorizedErrorf(err, "認証されていません。 invalid token claims")
			}

			auths, err := auth.NewAuth(token)
			if err != nil {
				return ferrors.NewUnauthorizedErrorf(err, "認証されていません。")
			}

			auth.SetAuth(cctx, auths)

			if err = next(cctx); err != nil {
				return xerrors.Errorf("next(): %w", err)
			}

			return nil
		}
	}
}

func getToken(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", xerrors.New("missing Authorization header")
	}

	if !strings.HasPrefix(authHeader, "Bearer ") {
		return "", xerrors.New("invalid Authorization header format")
	}

	return strings.TrimPrefix(authHeader, "Bearer "), nil
}

func getPublicKey(ctx context.Context, kid string) (publicKey any, err error) {
	var keySet jwk.Set
	if keySet, err = auth.GetKeySet(ctx); err != nil {
		return nil, xerrors.Errorf("unable to get key set. Error: %w", err)
	}

	key, ok := keySet.LookupKeyID(kid)
	if !ok {
		return nil, xerrors.Errorf("key not found in key set")
	}

	if err = key.Raw(&publicKey); err != nil {
		return nil, xerrors.Errorf("unable to get public key. Error: %w", err)
	}

	return publicKey, nil
}
