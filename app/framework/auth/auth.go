package auth

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	fcontext "github.com/y-nosuke/sample-task-api-go/app/framework/context"
	"golang.org/x/exp/slices"
	"golang.org/x/xerrors"
	"strings"
)

type ctxKey int

const (
	AUTH ctxKey = iota
)

func SetAuth(cctx fcontext.CustomContext, auth *Authentication) {
	cctx.WithValue(AUTH, auth)
}

func GetAuth(ctx context.Context) *Authentication {
	return ctx.Value(AUTH).(*Authentication)
}

type Authentication struct {
	UserId      uuid.UUID
	GivenName   string
	FamilyName  string
	Email       string
	Roles       []string
	Authorities []string
}

func NewAuth(token *jwt.Token) (*Authentication, error) {
	mapClaims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	iRoles := mapClaims["realm_access"].(map[string]interface{})["roles"].([]interface{})
	var roles []string
	for _, role := range iRoles {
		roles = append(roles, role.(string))
	}

	sUserId := mapClaims["sub"].(string)
	userId, err := uuid.Parse(sUserId)
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}

	return &Authentication{
		UserId:      userId,
		GivenName:   mapClaims["given_name"].(string),
		FamilyName:  mapClaims["family_name"].(string),
		Email:       mapClaims["email"].(string),
		Roles:       roles,
		Authorities: strings.Split(mapClaims["scope"].(string), " "),
	}, nil
}

func (a *Authentication) HasAuthority(authority string) bool {
	return slices.Contains(a.Authorities, authority)
}
