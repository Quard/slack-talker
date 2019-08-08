package rest_api

import (
	"log"
	"net/http"

	"github.com/Quard/slack-talker/internal/slack"
	"github.com/go-chi/chi"
	"github.com/go-pkgz/rest"
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
	return RestAPIServer{opts: opts, slackWebhookProcessor: slackWebhook}
}

func (srv RestAPIServer) Run() {
	router := srv.getRouter()

	log.Printf("REST API server listen on: %s", srv.opts.bind)
	log.Fatal(http.ListenAndServe(srv.opts.bind, router))
}

func (srv RestAPIServer) getRouter() chi.Router {
	router := chi.NewRouter()
	router.Route("/api/v1/slack", func(r chi.Router) {
		r.Get("/auth/code/", srv.slackOAuth2Code)
		r.Get("/auth/", srv.slackOAuth2Authorize)
		r.Post("/", srv.slackWebhookProcessor.Process)

	})

	return router
}

func responseError(w http.ResponseWriter, r *http.Request, err error) {
	w.WriteHeader(http.StatusBadRequest)
	rest.RenderJSON(w, r, rest.JSON{"error": err.Error()})
}
