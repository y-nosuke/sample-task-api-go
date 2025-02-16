package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/lestrrat-go/jwx/jwk"
	"golang.org/x/xerrors"
)

var Cfg Config

var (
	keySetCache jwk.Set
	cacheMutex  sync.Mutex
	cacheExpiry time.Time
)

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
	if time.Now().Before(cacheExpiry) && keySetCache != nil {
		return keySetCache, nil
	}

	fmt.Println("keySet is nil. fetching jwk set...")

	cacheMutex.Lock()
	defer cacheMutex.Unlock()

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

	if keySetCache, err = jwk.Fetch(ctx, jwksURI); err != nil {
		return nil, xerrors.Errorf("jwk fetch error: %w", err)
	}

	cacheExpiry = time.Now().Add(Cfg.CacheExpiry * time.Minute)

	return keySetCache, nil
}

type OIDCConfig struct {
	JwksURI string `json:"jwks_uri"`
}

func fetchJWKSURI(discoveryURL string) (jwksURI string, err error) {
	client := &http.Client{Timeout: 5 * time.Second}
	req, err := http.NewRequest(http.MethodGet, discoveryURL, nil)
	if err != nil {
		return "", err
	}

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		if closeErr := Body.Close(); closeErr != nil {
			err = xerrors.Errorf("original error: %v, defer close error: %w", err, closeErr)
			return
		}
	}(res.Body)

	var config OIDCConfig
	if err = json.NewDecoder(res.Body).Decode(&config); err != nil {
		return "", err
	}

	return config.JwksURI, nil
}
