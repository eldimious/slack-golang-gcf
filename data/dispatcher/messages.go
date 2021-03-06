package dispatcher

import (
	"github.com/ashwanthkumar/slack-go-webhook"
	"github.com/eldimious/slack-golang-gcf/config"
	domain "github.com/eldimious/slack-golang-gcf/domain/messages"
)

// Dispatcher sends notification to slack
type Dispatcher struct {
	config *config.Slack
}

// New initializes a Dispatcher
func New(config *config.Slack) *Dispatcher {
	return &Dispatcher{
		config: config,
	}
}

// SendMessage sends message to a channel
func (ds *Dispatcher) SendMessage(message *domain.Message) error {
	payload := slack.Payload{
		Text:      message.Text,
		Username:  message.Username,
		Channel:   message.Channel,
		IconEmoji: message.IconEmoji,
	}
	err := slack.Send(ds.config.WebhookUrl, "", payload)
	if len(err) > 0 {
		return err[0]
	}
	return nil
}

// SendSuccessMessage sends message to a channel with default iconemoji
func (ds *Dispatcher) SendSuccessMessage(message *domain.Message) error {
	payload := slack.Payload{
		Text:      message.Text,
		Username:  message.Username,
		Channel:   message.Channel,
		IconEmoji: ":heavy_check_mark:",
	}
	err := slack.Send(ds.config.WebhookUrl, "", payload)
	if len(err) > 0 {
		return err[0]
	}
	return nil
}

// SendErrorMessage sends message to a channel with default iconemoji
func (ds *Dispatcher) SendErrorMessage(message *domain.Message) error {
	payload := slack.Payload{
		Text:      message.Text,
		Username:  message.Username,
		Channel:   message.Channel,
		IconEmoji: ":bomb:",
	}
	err := slack.Send(ds.config.WebhookUrl, "", payload)
	if len(err) > 0 {
		return err[0]
	}
	return nil
}
