package auth

import "time"

type Config struct {
	DiscoveryUrl string        `env:"AUTH_DISCOVERY_URL"`
	JwksUrl      string        `env:"AUTH_JWKS_URL"`
	CacheExpiry  time.Duration `env:"AUTH_CACHE_EXPIRY" default:"10"`
}
