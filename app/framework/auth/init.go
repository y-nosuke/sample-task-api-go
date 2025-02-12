package auth

import (
	"context"
	"fmt"

	"github.com/caarlos0/env/v11"
	"github.com/lestrrat-go/jwx/jwk"
	"golang.org/x/xerrors"
)

var Cfg Config

var keySet jwk.Set

func init() {
	if err := env.Parse(&Cfg); err != nil {
		panic(fmt.Sprintf("parse config error: %v", err))
	}

	var err error
	if _, err = GetKeySet(context.Background()); err != nil {
		fmt.Printf("fetch jwk error: %v", err)
		return
	}

	fmt.Printf("auth initialized successfully.\n")
}

func GetKeySet(ctx context.Context) (jwk.Set, error) {
	if keySet != nil {
		return keySet, nil
	}

	fmt.Println("keySet is nil. fetching jwk set...")

	var err error
	if keySet, err = jwk.Fetch(ctx, Cfg.JwksUrl); err != nil {
		return nil, xerrors.Errorf("jwk fetch error: %w", err)
	}

	return keySet, nil
}
