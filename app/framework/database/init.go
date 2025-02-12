package database

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

var Cfg Config

func init() {
	if err := env.Parse(&Cfg); err != nil {
		panic(fmt.Sprintf("parse config error: %v", err))
	}

	con, err := sql.Open("mysql", Cfg.Dsn())
	if err != nil {
		panic(fmt.Sprintf("open database error: %v", err))
	}
	con.SetMaxIdleConns(Cfg.MaxIdleConn)
	con.SetMaxOpenConns(Cfg.MaxOpenConn)
	con.SetConnMaxLifetime(Cfg.ConnMaxLife * time.Second)

	boil.SetDB(con)
	boil.DebugMode = Cfg.Debug

	fmt.Printf("database initialized successfully.\n")
}
