package interfaces

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/sqlboiler/v4/boil"
	fcontext "github.com/y-nosuke/sample-task-api-go/framework/context/interfaces"
	"golang.org/x/xerrors"
)

type ctxKey int

const (
	TRANSACTION ctxKey = iota
)

func init() {
	fmt.Println("init database.")

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database_name := os.Getenv("DB_DATABASE_NAME")

	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + database_name + "?charset=utf8mb4&parseTime=True&loc=Local"
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
		fmt.Println("トランザクションを開始します。")
		cctx := fcontext.Cctx(ectx)

		tx, err := boil.BeginTx(cctx.Ctx, nil)
		if err != nil {
			return xerrors.Errorf(": %w", err)
		}
		defer func() {
			if p := recover(); p != nil {
				fmt.Println("ロールバックします。")
				tx.Rollback()
				panic(p)
			} else if err != nil {
				fmt.Println("ロールバックします。")
				tx.Rollback()
			} else {
				fmt.Println("コミットします。")
				err = tx.Commit()
			}
		}()
		cctx.WithValue(TRANSACTION, tx)

		err = next(ectx)

		return err
	}
}
