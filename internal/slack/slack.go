package slack

type Slack struct {
	authToken    string
	ClientID     string
	ClientSecret string
}

func NewSlack(authToken, clientID, clientSecret string) Slack {
	return Slack{
		authToken:    authToken,
		ClientID:     clientID,
		ClientSecret: clientSecret,
	}
}
