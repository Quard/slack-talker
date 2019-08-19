package slack

type Message struct {
	AuthToken string `json:"token"`
	Channel   string `json:"channel"`
	Text      string `json:"text"`
	ThreadTS  string `json:"thread_ts"`
}
