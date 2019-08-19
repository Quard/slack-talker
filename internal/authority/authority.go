package authority

import (
	context "context"
	"errors"

	"github.com/getsentry/sentry-go"
	grpc "google.golang.org/grpc"
)

type Authority struct {
	uri string
}

func NewAuthority(uri string) Authority {
	return Authority{uri: uri}
}

func (a Authority) GetUserByAuthToken() (*User, error) {
	conn, errConn := grpc.Dial(a.uri, grpc.WithInsecure())
	if errConn != nil {
		sentry.CaptureException(errConn)
		return nil, errConn
	}
	defer conn.Close()

	cl := NewInternalAPIClient(conn)
	user, err := cl.GetUserByAuthToken(
		context.Background(),
		&AuthToken{AuthToken: "0fe07a7cfc01bfc499946626a9cba5e09a31a75953e3cb552e947d718b454c5f"},
	)
	if err != nil {
		sentry.CaptureException(err)
		return nil, err
	}

	return user, nil
}

func (a Authority) GetUserBySlackID(slackID string) (*User, error) {
	conn, errConn := grpc.Dial(a.uri, grpc.WithInsecure())
	if errConn != nil {
		sentry.CaptureException(errConn)
		return nil, errConn
	}
	defer conn.Close()

	cl := NewInternalAPIClient(conn)
	user, err := cl.GetUserByProp(
		context.Background(),
		&UserProp{Name: "slack-user-id", Value: slackID},
	)
	if err != nil {
		sentry.CaptureException(err)
		return nil, err
	}

	return user, nil
}

func (a Authority) SetUserProps(authToken string, props []UserProp) error {
	conn, errConn := grpc.Dial(a.uri, grpc.WithInsecure())
	if errConn != nil {
		sentry.CaptureException(errConn)
		return errConn
	}
	defer conn.Close()

	cl := NewInternalAPIClient(conn)
	userProps := &UserProps{
		User:  &UserProps_AuthToken{AuthToken: authToken},
		Props: []*UserProp{},
	}
	for _, prop := range props {
		userProp := &UserProp{Name: prop.Name, Value: prop.Value}
		userProps.Props = append(userProps.Props, userProp)
	}
	intError, err := cl.SetUserProps(context.Background(), userProps)
	if err != nil {
		sentry.CaptureException(err)
		return err
	}
	if !intError.IsOk {
		return errors.New(intError.Msg)
	}

	return nil
}
