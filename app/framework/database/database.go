package database

import (
	"database/sql"

	fcontext "github.com/y-nosuke/sample-task-api-go/app/framework/context"
)

const (
	Transaction = "database.Transaction"
)

func SetTransaction(cctx fcontext.Context, tx *sql.Tx) {
	cctx.Set(Transaction, tx)
}

func GetTransaction(cctx fcontext.Context) *sql.Tx {
	return cctx.Get(Transaction).(*sql.Tx)
}
