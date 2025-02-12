package slack

type Config struct {
	SlackToken string `env:"SLACK_TOKEN"`
	ChannelID  string `env:"CHANNEL_ID"`
}
