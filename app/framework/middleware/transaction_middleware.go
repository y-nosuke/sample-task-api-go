package middleware

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/y-nosuke/sample-task-api-go/app/framework/context"
	"github.com/y-nosuke/sample-task-api-go/app/framework/database"
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
		fmt.Println("トランザクションを開始します。")

		tx, err := boil.BeginTx(cctx.GetContext(), nil)
		if err != nil {
			return xerrors.Errorf("boil.BeginTx(): %w", err)
		}
		defer func() {
			if p := recover(); p != nil {
				fmt.Println("トランザクションをロールバックします。")
				if rollbackErr := tx.Rollback(); rollbackErr != nil {
					err = xerrors.Errorf("original error: %v, defer rollback error: %v", err, rollbackErr)
				}
			} else if err != nil {
				fmt.Println("トランザクションをロールバックします。")
				if rollbackErr := tx.Rollback(); rollbackErr != nil {
					err = xerrors.Errorf("original error: %v, defer rollback error: %v", err, rollbackErr)
				}
			} else {
				fmt.Println("トランザクションをコミットします。")
				err = tx.Commit()
			}
		}()
		database.SetTransaction(cctx, tx)

		err = next(ectx)

		return
	}
}
