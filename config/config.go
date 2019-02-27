package config

import "github.com/eldimious/slack-golang-gcf/utils/env"

// Config is a struct that contains configuration variables
type Config struct {
	Slack *Slack
}

// Database is a struct that contains DB's configuration variables
type Slack struct {
	WebhookUrl string
}

// NewConfig creates a new Config struct
func NewConfig() (*Config, error) {
	env.CheckDotEnv()
	return &Config{
		Slack: &Slack{
			WebhookUrl: env.MustGet("WEBHOOK_URL"),
		},
	}, nil
}
