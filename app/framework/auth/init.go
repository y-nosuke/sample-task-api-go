package auth

import (
	"context"
	"fmt"

	"github.com/caarlos0/env/v11"
	"github.com/lestrrat-go/jwx/jwk"
)

var Cfg Config

var KeySet jwk.Set

func init() {
	if err := env.Parse(&Cfg); err != nil {
		panic(fmt.Sprintf("parse config error: %v", err))
	}

	var err error
	if KeySet, err = jwk.Fetch(context.Background(), Cfg.JwksUrl); err != nil {
		panic(fmt.Sprintf("fetch jwk error: %v", err))
	}

	fmt.Printf("auth initialized successfully.\n")
}
