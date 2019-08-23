package rest_api

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/Quard/slack-talker/internal/authority"
	"github.com/getsentry/sentry-go"
)

var regexStartCommand = regexp.MustCompile(`^(start)\s+(?:(\d{1,2}:\d{2})|([-\dhm]+))?\s*(.+)$`)
var regexTask = regexp.MustCompile(`\s*(\w+#\d+)\s*`)

func (srv RestAPIServer) slackTMCommand(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, authErr := srv.authority.GetUserBySlackID(r.PostFormValue("user_id"))
	if authErr != nil {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("unable to authorize"))
		return
	}

	text := r.PostFormValue("text")
	if parts := regexStartCommand.FindStringSubmatch(text); parts != nil {
		srv.slackTMCommandAdd(w, user, parts[1:])
	} else if strings.Trim(strings.ToLower(text), " ") == "done" {
		srv.slackTMCommandDone(w, user)
	} else if strings.Trim(strings.ToLower(text), " ")[:4] == "list" {
		fields := strings.Fields(text)
		if len(fields) > 1 {
			srv.slackTMCommandList(w, user, fields[1])
		} else {
			srv.slackTMCommandList(w, user, "")
		}
	}
}

func (srv RestAPIServer) slackTMCommandAdd(w http.ResponseWriter, user *authority.User, parts []string) {
	var task string

	if taskMatch := regexTask.FindStringSubmatch(parts[3]); taskMatch != nil {
		task = taskMatch[1]
		parts[3] = regexTask.ReplaceAllString(parts[3], "")
	}

	timeStart := time.Now()
	if parts[1] != "" {
		if start, err := time.Parse("15:04", parts[1]); err != nil {
			timeStart = start
		}
	} else if parts[2] != "" {
		if duration, err := time.ParseDuration(parts[2]); err != nil {
			timeStart = time.Now()
			timeStart = timeStart.Add(duration)
		}
	}

	record, err := srv.timekeeper.Add(user.ID, task, timeStart, parts[3])
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	definition := record.Task
	if definition == "" {
		definition = record.Comment
	}
	w.Write([]byte(
		fmt.Sprintf(
			"+ %s started at %s",
			definition,
			time.Unix(record.TimeStart, 0).Format("15:04"),
		),
	))

}

func (srv RestAPIServer) slackTMCommandDone(w http.ResponseWriter, user *authority.User) {
	record, err := srv.timekeeper.Finish(user.ID)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	definition := record.Task
	if definition == "" {
		definition = record.Comment
	}
	w.Write([]byte(
		fmt.Sprintf("%s %s", definition, record.GetDuration().String()),
	))
}

func (srv RestAPIServer) slackTMCommandList(w http.ResponseWriter, user *authority.User, dateStr string) {
	date := time.Now()
	if dateStr != "" {
		var err error
		date, err = time.Parse("2006-01-02", dateStr)
		if err != nil {
			sentry.CaptureException(err)
			return
		}
	}
	records, err := srv.timekeeper.GetGroupedForDate(user.ID, date)
	if err != nil {
		sentry.CaptureException(err)
		return
	}

	var total int64
	for _, record := range records {
		hour, minute := minutesToHouMinute(record.SpentMinutes)
		w.Write([]byte(fmt.Sprintf("%s - %dh %dm (%s)\n", record.Task, hour, minute, record.Comments)))
		total = total + record.SpentMinutes
	}
	hour, minute := minutesToHouMinute(total)
	w.Write([]byte(fmt.Sprintf("total: %dh %dm", hour, minute)))
}

func minutesToHouMinute(minutes int64) (int64, int64) {
	hours := minutes / 60
	return hours, minutes - hours*60
}
