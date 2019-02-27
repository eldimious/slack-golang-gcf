package messages

import (
	"fmt"

	"github.com/ashwanthkumar/slack-go-webhook"
	domain "github.com/eldimious/slack-golang-gcf/domain/messages"
)

type SlackConfig struct {
	webhookUrl string
}

// New initializes a KeyGetter
func New(webhookUrl string) *SlackConfig {
	return &SlackConfig{
		webhookUrl: webhookUrl,
	}
}

// GetPrivateKey reads the private key from the filesystem
func (kg *SlackConfig) SendNotification(message *domain.Message) (string, error) {
	payload := slack.Payload{
		Text:      message.Text,
		Username:  message.Username,
		Channel:   message.Channel,
		IconEmoji: message.IconEmoji,
	}
	err := slack.Send(kg.webhookUrl, "", payload)
	if len(err) > 0 {
		fmt.Printf("error: %s\n", err)
	}
	return "success", nil
}
