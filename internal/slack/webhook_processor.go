package slack

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/getsentry/sentry-go"
)

type WebhookProcessor struct {
	Slack Slack
}

func NewWebhookProcessor(_slack Slack) WebhookProcessor {
	return WebhookProcessor{Slack: _slack}
}

func (p WebhookProcessor) Process(w http.ResponseWriter, r *http.Request) {
	var webhookEvent WebhookEvent
	if err := json.NewDecoder(r.Body).Decode(&webhookEvent); err != nil {
		sentry.CaptureException(err)
		return
	}

	switch webhookEvent.Type {
	case "url_verification":
		p.processUrlVerification(webhookEvent, w, r)
	case "event_callback":
		switch webhookEvent.Event.Type {
		case "message":
			p.processMessage(webhookEvent, w, r)
		}
	}

}

func (p WebhookProcessor) processUrlVerification(webhookEvent WebhookEvent, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp := struct {
		Challenge string `json:"challenge"`
	}{Challenge: webhookEvent.Challenge}
	json.NewEncoder(w).Encode(resp)
}

func (p WebhookProcessor) processMessage(webhookEvent WebhookEvent, w http.ResponseWriter, r *http.Request) {
	log.Printf("[new MSG] <%s> %s\n", webhookEvent.Event.User, webhookEvent.Event.Text)
}
