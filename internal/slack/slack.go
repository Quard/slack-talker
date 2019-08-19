package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/getsentry/sentry-go"
)

const host = "https://slack.com/api/"

type Slack struct {
	authToken    string
	ClientID     string
	ClientSecret string
}

func NewSlack(authToken, clientID, clientSecret string) Slack {
	return Slack{
		authToken:    authToken,
		ClientID:     clientID,
		ClientSecret: clientSecret,
	}
}

func (s Slack) NewMessage(channel, text string) Message {
	return Message{
		AuthToken: s.authToken,
		Channel:   channel,
		Text:      text,
	}
}

func (s Slack) SendMessage(msg Message) error {
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(msg); err != nil {
		sentry.CaptureException(err)
		return err
	}

	request, err := http.NewRequest("POST", host+"chat.postMessage", buf)
	if err != nil {
		sentry.CaptureException(err)
		return err
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.authToken))

	client := http.Client{}
	if _, err := client.Do(request); err != nil {
		sentry.CaptureException(err)
		return err
	}

	return nil
}
