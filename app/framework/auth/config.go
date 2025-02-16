package auth

type Config struct {
	DiscoveryUrl string `env:"AUTH_DISCOVERY_URL"`
	JwksUrl      string `env:"AUTH_JWKS_URL"`
}
