package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

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

	jwksURI := Cfg.JwksUrl
	var err error

	if jwksURI == "" {
		if Cfg.DiscoveryUrl == "" {
			return nil, xerrors.New("jwks url is required")
		}

		jwksURI, err = fetchJWKSURI(Cfg.DiscoveryUrl)
		if err != nil {
			return nil, xerrors.Errorf("failed to fetch JWKS URI: %w", err)
		}
	}

	if keySet, err = jwk.Fetch(ctx, jwksURI); err != nil {
		return nil, xerrors.Errorf("jwk fetch error: %w", err)
	}

	return keySet, nil
}

type OIDCConfig struct {
	JwksURI string `json:"jwks_uri"`
}

func fetchJWKSURI(discoveryURL string) (string, error) {
	client := &http.Client{Timeout: 5 * time.Second}
	req, err := http.NewRequest(http.MethodGet, discoveryURL, nil)
	if err != nil {
		return "", err
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {

		}
	}(resp.Body)

	var config OIDCConfig
	if err := json.NewDecoder(resp.Body).Decode(&config); err != nil {
		return "", err
	}

	return config.JwksURI, nil
}
