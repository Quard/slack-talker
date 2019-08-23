package rest_api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/Quard/slack-talker/internal/authority"
	"github.com/Quard/slack-talker/internal/poindexter"
)

func (srv RestAPIServer) slackRLCommand(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fields := strings.Fields(r.PostFormValue("text"))

	user, authErr := srv.authority.GetUserBySlackID(r.PostFormValue("user_id"))
	if authErr != nil {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("unable to authorize"))
		return
	}

	switch strings.ToLower(fields[0]) {
	case "add":
		srv.slackRLCommandAdd(w, user, fields[1])
	case "list":
		srv.slackRLCommandList(w, user)
	case "read":
		srv.slackRLCommandMarkAsRead(w, user, fields[1])
	}
}

func (srv RestAPIServer) slackRLCommandAdd(w http.ResponseWriter, user *authority.User, url string) {
	record, err := srv.poindexter.AddReadingListRecord(user, url)
	if err != nil {
		log.Println("ERR", err)
	}

	w.Write([]byte(fmt.Sprintf("'%s' added", record.Title)))
}

type readinglistRecordText struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type readinglistRecordAccessory struct {
	Type     string `json:"type"`
	ImageURL string `json:"image_url"`
	AltText  string `json:"alt_text"`
}

type readinglistRecordsField struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type readinglistRecordsList struct {
	Type      string                     `json:"type"`
	Text      readinglistRecordText      `json:"text"`
	Accessory readinglistRecordAccessory `json:"accessory"`
	Fields    []readinglistRecordsField  `json:"fields"`
}

func (srv RestAPIServer) slackRLCommandList(w http.ResponseWriter, user *authority.User) {
	ch := make(chan *poindexter.ReadingListRecord)
	err := srv.poindexter.ListReadingListRecord(user, ch)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	links := []readinglistRecordsList{}
	for record := range ch {
		link, err := url.Parse(record.Url)
		if err != nil {
			continue
		}
		titleTmpl := "*<%s|%s>*"
		if record.IsRead {
			titleTmpl = "~*<%s|%s>*~"
		}
		links = append(
			links,
			readinglistRecordsList{
				Type: "section",
				Text: readinglistRecordText{
					Type: "mrkdwn",
					Text: fmt.Sprintf(titleTmpl, record.Url, record.Title),
				},
				Accessory: readinglistRecordAccessory{
					Type:     "image",
					ImageURL: record.ImageUrl,
					AltText:  record.Title,
				},
				Fields: []readinglistRecordsField{
					{Type: "mrkdwn", Text: "*ID*"},
					{Type: "mrkdwn", Text: "*source*"},
					// {Type: "mrkdwn", Text: "*added*"},
					{Type: "plain_text", Text: record.Id},
					{Type: "mrkdwn", Text: link.Hostname()},
					// {Type: "mrkdwn", Text: time.Unix(record.Created, 0).String()},
				},
			},
		)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(
		struct {
			Blocks []readinglistRecordsList `json:"blocks"`
		}{
			Blocks: links,
		},
	)
}

func (srv RestAPIServer) slackRLCommandMarkAsRead(w http.ResponseWriter, user *authority.User, ID string) {
	record, err := srv.poindexter.MarkRecordAsRead(user, ID)
	if err != nil {
		w.Write([]byte("unable to update record"))
	} else {
		w.Write([]byte(fmt.Sprintf("'%s' marked as read", record.Title)))
	}
}
