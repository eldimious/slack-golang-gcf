package messageValidator

import (
	"net/url"
	"strings"
)

type MessageValidator struct {
	Text      string `binding:"required" json:"text"`
	Username  string `binding:"required" json:"userName"`
	Channel   string `binding:"required" json:"channel"`
	IconEmoji string `json:"iconEmoji"`
	Type      string `json:"type"`
}

func Validate(a *MessageValidator) url.Values {
	errs := url.Values{}

	// check if the title empty
	if a.Text == "" {
		errs.Add("text", "The text field is required!")
	}
	if a.Username == "" {
		errs.Add("userName", "The userName field is required!")
	}
	if a.Channel == "" {
		errs.Add("channel", "The channel field is required!")
	}

	if a.Channel != "" && strings.HasPrefix(a.Channel, "#") != true {
		errs.Add("channel", "The channel field must start with #")
	}

	return errs
}
