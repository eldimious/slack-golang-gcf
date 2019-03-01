package function

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/eldimious/slack-golang-gcf/config"
	dispatcher "github.com/eldimious/slack-golang-gcf/data/dispatcher"
	messages "github.com/eldimious/slack-golang-gcf/domain/messages"
	validator "github.com/eldimious/slack-golang-gcf/router"
)

// SendNotification - Send notification to slack
func SendNotification(w http.ResponseWriter, r *http.Request) {
	data := &validator.MessageValidator{}
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.Header().Set("Content-type", "applciation/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		log.Println(err.Error())
		return
	}
	validationErrs := validator.Validate(data)
	if len(validationErrs) > 0 {
		err := map[string]interface{}{"validationError": validationErrs}
		w.Header().Set("Content-type", "applciation/json")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error: %s", err)
		json.NewEncoder(w).Encode(err)
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
		log.Println(err.Error())
		w.Header().Set("Content-type", "applciation/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	messagesDispatcher := dispatcher.New(configuration.Slack)
	messagesSvc := messages.NewService(messagesDispatcher)
	dispatcherError := messagesSvc.SendMessage(message)
	if dispatcherError != nil {
		w.Header().Set("Content-type", "applciation/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(dispatcherError.Error()))
		log.Println(dispatcherError.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "applciation/json")
	w.Write([]byte("Message was sent!"))
	return
}
