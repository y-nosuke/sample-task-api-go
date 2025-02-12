package auth

type Config struct {
	JwksUrl string `env:"AUTH_JWKS_URL"`
}
