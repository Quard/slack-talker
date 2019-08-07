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
	Bind           string `short:"b" long:"bind" env:"BIND" default:"localhost:5005" description:"address:port to listen"`
	SlackAuthToken string `long:"slack-token" env:"SLACK_AUTHTOKEN"`
	SentryDSN      string `long:"sentry-dsn" env:"SENTRY_DSN"`
}

func main() {
	parser := flags.NewParser(&opts, flags.PrintErrors|flags.PassDoubleDash)
	if _, err := parser.Parse(); err != nil {
		log.Fatal(err)
	}

	sentry.Init(sentry.ClientOptions{Dsn: opts.SentryDSN})
	defer sentry.Flush(time.Second * 5)

	slackWebhookProcessor := slack.NewWebhookProcessor(slack.NewSlack(opts.SlackAuthToken))
	srv := rest_api.NewRestAPIServer(opts.Bind, slackWebhookProcessor)
	srv.Run()
}
