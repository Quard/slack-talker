package slack

type Slack struct {
	authToken string
}

func NewSlack(authToken string) Slack {
	return Slack{authToken: authToken}
}
