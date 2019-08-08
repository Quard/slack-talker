package rest_api

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"

	"github.com/getsentry/sentry-go"

	"github.com/Quard/slack-talker/internal/authority"
)

type oAuth2AccessResponse struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	UserID      string `json:"user_id"`
}

func (srv RestAPIServer) slackOAuth2Authorize(w http.ResponseWriter, r *http.Request) {
	params := url.Values{}
	params.Add("client_id", srv.slackWebhookProcessor.Slack.ClientID)
	params.Add("scope", "channels:history")
	params.Add("scope", "group:history")
	params.Add("scope", "reactions:read")

	w.Header().Set("Location", "https://slack.com/oauth/authorize?"+params.Encode())
	w.WriteHeader(http.StatusMovedPermanently)
}

func (srv RestAPIServer) slackOAuth2Code(w http.ResponseWriter, r *http.Request) {
	respData, err := slackOAuth2GetAccessToken(
		srv.slackWebhookProcessor.Slack.ClientID,
		srv.slackWebhookProcessor.Slack.ClientSecret,
		r.URL.Query().Get("code"),
	)
	if err != nil {
		responseError(w, r, err)
		return
	}

	authToken, errCookie := r.Cookie("t")
	if errCookie != nil {
		responseError(w, r, errors.New("unauthorized user"))
	}
	err = authority.SetUserProps(
		authToken.Value,
		[]authority.UserProp{
			{Name: "slack-access-token", Value: respData.AccessToken},
			{Name: "slack-user-id", Value: respData.UserID},
		},
	)
	if err != nil {
		responseError(w, r, err)
	}
}

func slackOAuth2GetAccessToken(clientID, clientSecret, code string) (oAuth2AccessResponse, error) {
	var respData oAuth2AccessResponse

	buf := new(bytes.Buffer)
	buf.WriteString(url.Values{"code": []string{code}}.Encode())
	r, err := http.NewRequest("POST", "https://slack.com/api/oauth.access", buf)
	if err != nil {
		sentry.CaptureException(err)
		return respData, err
	}

	r.SetBasicAuth(clientID, clientSecret)
	r.Header.Set("content-type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, errReq := client.Do(r)
	if errReq != nil {
		sentry.CaptureException(err)
		return respData, errReq
	}

	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		sentry.CaptureException(err)
		return respData, err
	}

	return respData, nil
}
