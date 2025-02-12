package mail

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

var Cfg Config

func init() {
	if err := env.Parse(&Cfg); err != nil {
		panic(fmt.Sprintf("parse config error: %v", err))
	}

	fmt.Printf("mail initialized successfully.\n")
}
