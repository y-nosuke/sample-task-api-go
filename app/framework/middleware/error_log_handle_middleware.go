package middleware

import (
	"errors"
	"fmt"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/y-nosuke/sample-task-api-go/app/framework/context"
	ferrors "github.com/y-nosuke/sample-task-api-go/app/framework/errors"
	"golang.org/x/xerrors"
)

func ErrorLogHandleMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ectx echo.Context) (err error) {
		cctx := context.CastContext(ectx)
		fmt.Println("ErrorLogHandleMiddleware start. エラーログハンドラーを実行します。")

		defer func() {
			if p := recover(); p != nil {
				ectx.Logger().Infof("panic occurred!: %+v", p)
				fmt.Printf("panic occurred!: %+v\n", p)
				panic(fmt.Sprintf("panic occurred. %v", p))
			} else if err != nil {
				if businessError, ok := ferrors.AsBusinessError(err); ok {
					ectx.Logger().Warnf("business error!: %+v, orinal error: %+v", businessError, businessError.OriginalError())
					fmt.Printf("business error!: %+v, orinal error: %+v\n", businessError, businessError.OriginalError())
				} else {
					var he *echo.HTTPError
					if errors.As(err, &he) {
						ectx.Logger().Infof("http error!: %+v", he)
						fmt.Printf("http error!: %+v\n", he)
					} else if errors.Is(err, io.ErrUnexpectedEOF) {
						ectx.Logger().Infof("unexpected EOF error!: %+v", he)
						fmt.Printf("unexpected EOF error!: %+v\n", he)
					} else if errors.Is(err, io.EOF) {
						ectx.Logger().Infof("EOF error!: %+v", he)
						fmt.Printf("EOF error!: %+v\n", he)
					} else {
						ectx.Logger().Errorf("system error!: %+v", err)
						fmt.Printf("system error!: %+v\n", err)
					}
				}
				err = xerrors.Errorf("next(): %w", err)
			}
		}()

		err = next(cctx)

		return
	}
}
