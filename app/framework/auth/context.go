package auth

import (
	"github.com/google/uuid"
	fcontext "github.com/y-nosuke/sample-task-api-go/app/framework/context"
)

const (
	Auth = "auth.Auth"
)

func SetAuth(cctx fcontext.Context, auth *Authentication) {
	cctx.Set(Auth, auth)
}

func GetAuth(cctx fcontext.Context) *Authentication {
	return cctx.Get(Auth).(*Authentication)
}

func GetUserId(cctx fcontext.Context) uuid.UUID {
	return cctx.Get(Auth).(*Authentication).UserId
}
