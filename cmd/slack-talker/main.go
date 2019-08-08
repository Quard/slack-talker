package main

import (
	"log"
	"time"

	"github.com/Quard/slack-talker/internal/rest_api"
	"github.com/Quard/slack-talker/internal/slack"
	"github.com/getsentry/sentry-go"
	"github.com/jessevdk/go-flags"
)

var opts struct {
	Bind              string `short:"b" long:"bind" env:"BIND" default:"localhost:5005" description:"address:port to listen"`
	SlackAuthToken    string `long:"slack-token" env:"SLACK_AUTHTOKEN"`
	SlackClientID     string `long:"slack-client-id" env:"SLACK_CLIENT_ID"`
	SlackClientSecret string `long:"slack-client-secret" env:"SLACK_CLIENT_SECRET"`
	SentryDSN         string `long:"sentry-dsn" env:"SENTRY_DSN"`
}

func main() {
	parser := flags.NewParser(&opts, flags.PrintErrors|flags.PassDoubleDash)
	if _, err := parser.Parse(); err != nil {
		log.Fatal(err)
	}

	sentry.Init(sentry.ClientOptions{Dsn: opts.SentryDSN})
	defer sentry.Flush(time.Second * 5)

	_slack := slack.NewSlack(
		opts.SlackAuthToken,
		opts.SlackClientID,
		opts.SlackClientSecret,
	)
	slackWebhookProcessor := slack.NewWebhookProcessor(_slack)
	srv := rest_api.NewRestAPIServer(opts.Bind, slackWebhookProcessor)
	srv.Run()
}
