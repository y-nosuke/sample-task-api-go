package middleware

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/y-nosuke/sample-task-api-go/app/framework/context"
	"github.com/y-nosuke/sample-task-api-go/app/framework/database"
	ferrors "github.com/y-nosuke/sample-task-api-go/app/framework/errors"
	"golang.org/x/xerrors"
	"os"
	"time"
)

func init() {
	fmt.Println("init repository.")

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	databaseName := os.Getenv("DB_DATABASE_NAME")

	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + databaseName + "?charset=utf8mb4&parseTime=True&loc=Local"
	con, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	con.SetMaxIdleConns(10)
	con.SetMaxOpenConns(10)
	con.SetConnMaxLifetime(300 * time.Second)

	boil.SetDB(con)
	boil.DebugMode = true
}

func TransactionMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ectx echo.Context) (err error) {
		cctx := context.CastContext(ectx)
		fmt.Println("TransactionMiddleware start. トランザクションを開始します。")

		tx, err := boil.BeginTx(cctx.GetContext(), nil)
		if err != nil {
			return xerrors.Errorf("boil.BeginTx(): %w", err)
		}

		fmt.Println("トランザクションを開始しました。")

		defer func() {
			if p := recover(); p != nil {
				if rollbackErr := tx.Rollback(); rollbackErr != nil {
					panic(fmt.Sprintf("original panic: %v, defer rollback error: %v", p, rollbackErr))
				}
				fmt.Printf("トランザクションをロールバックしました。 panic: %v\n", p)
				panic(fmt.Sprintf("panic occurred. %v", p))
			} else if err != nil {
				if businessError, ok := ferrors.AsBusinessError(err); ok {
					if businessError.DoRollBack() {
						if rollbackErr := tx.Rollback(); rollbackErr != nil {
							err = xerrors.Errorf("original error: %v, defer rollback error: %w", err, rollbackErr)
							return
						}
						fmt.Printf("トランザクションをロールバックしました。 business error: %v\n", p)
						err = xerrors.Errorf("next(): %w", err)
					} else {
						if commitErr := tx.Commit(); commitErr != nil {
							err = xerrors.Errorf("original error: %v, defer commit error: %w", p, commitErr)
							return
						}
						fmt.Println("トランザクションをコミットしました。")
						err = xerrors.Errorf("next(): %w", err)
					}
				} else {
					if rollbackErr := tx.Rollback(); rollbackErr != nil {
						err = xerrors.Errorf("original error: %v, defer rollback error: %w", err, rollbackErr)
						return
					}
					fmt.Printf("トランザクションをロールバックしました。 system error: %v\n", p)
					err = xerrors.Errorf("next(): %w", err)
				}
			} else {
				if commitErr := tx.Commit(); commitErr != nil {
					err = xerrors.Errorf("original error: %+v, defer commit error: %+w", err, commitErr)
				}
				fmt.Println("トランザクションをコミットしました。")
			}
		}()

		database.SetTransaction(cctx, tx)

		err = next(cctx)

		return
	}
}
