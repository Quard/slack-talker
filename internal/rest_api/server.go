package rest_api

import (
	"log"
	"net/http"

	"github.com/Quard/slack-talker/internal/slack"
	"github.com/go-chi/chi"
)

type Opts struct {
	bind string
}

type RestAPIServer struct {
	opts                  Opts
	slackWebhookProcessor slack.WebhookProcessor
}

func NewRestAPIServer(bind string, slackWebhook slack.WebhookProcessor) RestAPIServer {
	opts := Opts{bind: bind}
	return RestAPIServer{
		opts:                  opts,
		slackWebhookProcessor: slackWebhook,
	}
}

func (srv RestAPIServer) Run() {
	router := srv.getRouter()

	log.Printf("REST API server listen on: %s", srv.opts.bind)
	log.Fatal(http.ListenAndServe(srv.opts.bind, router))
}

func (srv RestAPIServer) getRouter() chi.Router {
	router := chi.NewRouter()
	router.Route("/api/v1/slack", func(r chi.Router) {
		r.Post("/", srv.slackWebhookProcessor.Process)
	})

	return router
}
