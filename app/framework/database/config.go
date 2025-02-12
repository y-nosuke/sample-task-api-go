package database

import (
	"fmt"
	"time"
)

type Config struct {
	Host        string        `env:"DB_HOST"`
	Port        string        `env:"DB_PORT"`
	User        string        `env:"DB_USER"`
	Password    string        `env:"DB_PASSWORD"`
	Database    string        `env:"DB_DATABASE_NAME"`
	Debug       bool          `env:"DB_DEBUG" default:"false"`
	MaxIdleConn int           `env:"DB_MAX_IDLE_CONN" default:"10"`
	MaxOpenConn int           `env:"DB_MAX_OPEN_CONN" default:"10"`
	ConnMaxLife time.Duration `env:"DB_CONN_MAX_LIFE" default:"300"`
}

func (c Config) Dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", c.User, c.Password, c.Host, c.Port, c.Database)
}
