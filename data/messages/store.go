package messages

import (
	"fmt"

	"github.com/ashwanthkumar/slack-go-webhook"
	"github.com/eldimious/slack-golang-gcf/config"
	domain "github.com/eldimious/slack-golang-gcf/domain/messages"
)

// Dispatcher sends notification to slack groups
type Dispatcher struct {
	config *config.Slack
}

// New initializes a Dispatcher
func New(config *config.Slack) *Dispatcher {
	return &Dispatcher{
		config: config,
	}
}

// SendMessage send message
func (kg *Dispatcher) SendMessage(message *domain.Message) (string, error) {
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
