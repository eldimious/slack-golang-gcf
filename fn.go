package function

import (
	"encoding/json"
	"net/http"

	"github.com/eldimious/slack-golang-gcf/config"
	messagesStore "github.com/eldimious/slack-golang-gcf/data/messages"
	messages "github.com/eldimious/slack-golang-gcf/domain/messages"
)

type messageValidator struct {
	Text      string
	Username  string
	Channel   string
	IconEmoji string
}

// GetUserDetails - Get one user's details from randomuser.me API
func GetUserDetails(w http.ResponseWriter, r *http.Request) {
	var data messageValidator
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		panic(err)
	}

	//data := decoder.Decode(&messages.Message)
	message := &messages.Message{
		Text:      data.Text,
		Username:  data.Username,
		Channel:   data.Channel,
		IconEmoji: data.IconEmoji,
	}

	configuration, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	messagesDispatcher := messagesStore.New(configuration.Slack)
	messagesSvc := messages.NewService(messagesDispatcher)
}
