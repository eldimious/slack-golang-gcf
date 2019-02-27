package messages

import (
	"fmt"

	"github.com/ashwanthkumar/slack-go-webhook"
	"github.com/eldimious/slack-golang-gcf/config"
	domain "github.com/eldimious/slack-golang-gcf/domain/messages"
)

type SlackConfig struct {
	config *config.Slack
}

// New initializes a KeyGetter
func New(config *config.Slack) *SlackConfig {
	return &SlackConfig{
		config: config,
	}
}

// GetPrivateKey reads the private key from the filesystem
func (kg *SlackConfig) SendMessage(message *domain.Message) (interface{}, error) {
	payload := slack.Payload{
		Text:      message.Text,
		Username:  message.Username,
		Channel:   message.Channel,
		IconEmoji: message.IconEmoji,
	}
	err := slack.Send(kg.config.WebhookUrl, "", payload)
	if len(err) > 0 {
		fmt.Printf("error: %s\n", err)
	}
	return "success", nil
}
