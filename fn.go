package function

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/ashwanthkumar/slack-go-webhook"
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
	webhookUrl := "https://hooks.slack.com/services/foo/bar/baz"
	payload := slack.Payload{
		Text:      "Hello from <https://github.com/ashwanthkumar/slack-go-webhook|slack-go-webhook>, a Go-Lang library to send slack webhook messages.\n<https://golangschool.com/wp-content/uploads/golang-teach.jpg|golang-img>",
		Username:  "robot",
		Channel:   "#general",
		IconEmoji: ":monkey_face:",
	}
	err := slack.Send(webhookUrl, "", payload)
	if len(err) > 0 {
		fmt.Printf("error: %s\n", err)
	}
	randomUserClient := http.Client{
		Timeout: time.Second * 3,
	}

	req, err := http.NewRequest(http.MethodGet, "https://randomuser.me/api/", nil)
	if err != nil {
		log.Fatal(err)
		return
	}

	res, err2 := randomUserClient.Do(req)
	if err2 != nil {
		log.Fatal(err2)
		return
	}

	body, err3 := ioutil.ReadAll(res.Body)
	if err3 != nil {
		log.Fatal(err3)
	}

	var o map[string]interface{}
	json.Unmarshal([]byte(body), &o)

	results := o["results"].([]interface{})
	result := results[0].(map[string]interface{})

	result["generator"] = "google-cloud-function"

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
