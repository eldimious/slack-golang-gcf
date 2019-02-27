package function

import (
	"net/http"

	"github.com/eldimious/slack-golang-gcf/config"
	messagesStore "github.com/eldimious/slack-golang-gcf/data/messages"
)

// GetUserDetails - Get one user's details from randomuser.me API
func GetUserDetails(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}
	v := r.Form
	owner := r.Form.Get("owner")
	name := r.Form.Get("name")
	configuration, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	messagesRepo := messagesStore.New(configuration.Slack.WebhookUrl)

}
