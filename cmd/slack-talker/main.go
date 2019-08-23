package main

import (
	"log"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/jessevdk/go-flags"

	"github.com/Quard/slack-talker/internal/authority"
	"github.com/Quard/slack-talker/internal/poindexter"
	"github.com/Quard/slack-talker/internal/rest_api"
	"github.com/Quard/slack-talker/internal/slack"
	"github.com/Quard/slack-talker/internal/timekeeper_rpc"
)

var opts struct {
	Bind              string `short:"b" long:"bind" env:"BIND" default:"localhost:5005" description:"address:port to listen"`
	AuthorityURI      string `long:"authority-uri" env:"AUTHORITY_URI"`
	PoindexterURI     string `long:"poindexter-uri" env:"POINDEXTER_URI"`
	TimeKeeperURI     string `long:"timekeeper-uri" env:"TIMEKEEPER_URI"`
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
	poindexterService := poindexter.NewPoindexter(opts.PoindexterURI)
	auhtorityService := authority.NewAuthority(opts.AuthorityURI)
	timerkeeperService := timekeeper_rpc.NewTimeKeeper(opts.TimeKeeperURI)
	slackWebhookProcessor := slack.NewWebhookProcessor(_slack)
	srv := rest_api.NewRestAPIServer(
		opts.Bind,
		slackWebhookProcessor,
		auhtorityService,
		poindexterService,
		timerkeeperService,
	)
	srv.Run()
}
