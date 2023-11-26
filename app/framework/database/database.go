package database

import (
	"context"
	"database/sql"

	fcontext "github.com/y-nosuke/sample-task-api-go/app/framework/context"
)

type ctxKey int

const (
	TRANSACTION ctxKey = iota
)

func SetTransaction(cctx fcontext.CustomContext, tx *sql.Tx) {
	cctx.WithValue(TRANSACTION, tx)
}

func GetTransaction(ctx context.Context) *sql.Tx {
	return ctx.Value(TRANSACTION).(*sql.Tx)
}
