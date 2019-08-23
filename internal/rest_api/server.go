package rest_api

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-pkgz/rest"
	"github.com/go-pkgz/rest/logger"

	"github.com/Quard/slack-talker/internal/authority"
	"github.com/Quard/slack-talker/internal/poindexter"
	"github.com/Quard/slack-talker/internal/slack"
	"github.com/Quard/slack-talker/internal/timekeeper_rpc"
)

type Opts struct {
	bind string
}

type RestAPIServer struct {
	opts                  Opts
	slackWebhookProcessor slack.WebhookProcessor
	authority             authority.Authority
	poindexter            poindexter.Poindexter
	timekeeper            timekeeper_rpc.TimeKeeper
}

func NewRestAPIServer(
	bind string,
	slackWebhook slack.WebhookProcessor,
	authoritySerive authority.Authority,
	poindexterService poindexter.Poindexter,
	timekeeperService timekeeper_rpc.TimeKeeper,
) RestAPIServer {
	opts := Opts{bind: bind}
	return RestAPIServer{
		opts:                  opts,
		slackWebhookProcessor: slackWebhook,
		authority:             authoritySerive,
		poindexter:            poindexterService,
		timekeeper:            timekeeperService,
	}
}

func (srv RestAPIServer) Run() {
	router := srv.getRouter()

	log.Printf("REST API server listen on: %s", srv.opts.bind)
	log.Fatal(http.ListenAndServe(srv.opts.bind, router))
}

func (srv RestAPIServer) getRouter() chi.Router {
	router := chi.NewRouter()
	router.Use(logger.Logger)
	router.Get("/", srv.healthCheck)
	router.Route("/api/v1/slack", func(r chi.Router) {
		r.Get("/auth/code/", srv.slackOAuth2Code)
		r.Get("/auth/", srv.slackOAuth2Authorize)
		r.Post("/cmd/rl/", srv.slackRLCommand)
		r.Post("/cmd/tm/", srv.slackTMCommand)
		r.Post("/", srv.slackWebhookProcessor.Process)

	})

	return router
}

func (srv RestAPIServer) healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func responseError(w http.ResponseWriter, r *http.Request, err error) {
	w.WriteHeader(http.StatusBadRequest)
	rest.RenderJSON(w, r, rest.JSON{"error": err.Error()})
}
