package mail

type Config struct {
	Host string `env:"MAIL_HOST"`
	Port int    `env:"MAIL_PORT"`
	From string `env:"MAIL_FROM"`
	To   string `env:"MAIL_TO"`
}
