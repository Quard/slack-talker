package slack

type Event struct {
	Type        string `json:"type"`
	SubType     string `json:"subtype"`
	ClientMsgID string `json:"client_msg_id"`
	User        string `json:"user"`
	ItemUser    string `json:"item_user"`
	Text        string `json:"text"`
	Reaction    string `json:"reaction"`
	Channel     string `json:"channel"`
	ChannleType string `json:"channel_type"`
	TS          string `json:"ts"`
}

type WebhookEvent struct {
	Type      string `json:"type"`
	SubType   string `json:"subtype"`
	Token     string `json:"token"`
	TeamID    string `json:"team_id"`
	Challenge string `json:"challenge"`
	Event     Event  `json:"event"`
}
