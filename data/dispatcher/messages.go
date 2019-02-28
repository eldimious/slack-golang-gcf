package dispatcher

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
func (ds *Dispatcher) SendMessage(message *domain.Message) (string, error) {
	payload := slack.Payload{
		Text:      message.Text,
		Username:  message.Username,
		Channel:   message.Channel,
		IconEmoji: message.IconEmoji,
	}
	err := slack.Send(ds.config.WebhookUrl, "", payload)
	if len(err) > 0 {
		fmt.Printf("error: %s\n", err)
	}
	return "Your message has been sent successfully.", nil
}

func (ds *Dispatcher) SendSuccessMessage(message *domain.Message) (string, error) {
	payload := slack.Payload{
		Text:      message.Text,
		Username:  message.Username,
		Channel:   message.Channel,
		IconEmoji: ":heavy_check_mark:",
	}
	err := slack.Send(ds.config.WebhookUrl, "", payload)
	if len(err) > 0 {
		fmt.Printf("error: %s\n", err)
	}
	return "Your message has been sent successfully.", nil
}

func (ds *Dispatcher) SendErrorMessage(message *domain.Message) (string, error) {
	payload := slack.Payload{
		Text:      message.Text,
		Username:  message.Username,
		Channel:   message.Channel,
		IconEmoji: ":bomb:",
	}
	err := slack.Send(ds.config.WebhookUrl, "", payload)
	if len(err) > 0 {
		fmt.Printf("error: %s\n", err)
	}
	return "Your message has been sent successfully.", nil
}
