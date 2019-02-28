package function

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/eldimious/slack-golang-gcf/config"
	dispatcher "github.com/eldimious/slack-golang-gcf/data/dispatcher"
	messages "github.com/eldimious/slack-golang-gcf/domain/messages"
)

type messageValidator struct {
	Text      string
	Username  string
	Channel   string
	IconEmoji string
	Type      string
}

// SendNotification - Send notification to slack group
func SendNotification(w http.ResponseWriter, r *http.Request) {
	var data messageValidator
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		log.Println(err.Error())
		return
	}
	message := &messages.Message{
		Text:      data.Text,
		Username:  data.Username,
		Channel:   data.Channel,
		IconEmoji: data.IconEmoji,
		Type:      data.Type,
	}

	configuration, err := config.NewConfig()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		log.Println(err.Error())
		return
	}
	messagesDispatcher := dispatcher.New(configuration.Slack)
	messagesSvc := messages.NewService(messagesDispatcher)
	_, dispatcherError := messagesSvc.SendMessage(message)
	if dispatcherError != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(dispatcherError.Error()))
		log.Println(dispatcherError.Error())
		return
	}
	// all good. write our message.
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Message was sent!"))
	return
}
