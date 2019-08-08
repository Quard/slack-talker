package authority

import (
	context "context"
	"errors"

	"github.com/getsentry/sentry-go"
	grpc "google.golang.org/grpc"
)

func GetUserByAuthToken() (*User, error) {
	conn, errConn := grpc.Dial("localhost:5001", grpc.WithInsecure())
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

type UserProp struct {
	Name  string
	Value string
}

func SetUserProps(authToken string, props []UserProp) error {
	conn, errConn := grpc.Dial("localhost:5001", grpc.WithInsecure())
	if errConn != nil {
		sentry.CaptureException(errConn)
		return errConn
	}
	defer conn.Close()

	cl := NewInternalAPIClient(conn)
	userProps := &UserProps{
		User:  &UserProps_AuthToken{AuthToken: authToken},
		Props: []*UserProps_Prop{},
	}
	for _, prop := range props {
		userProp := &UserProps_Prop{Name: prop.Name, Value: prop.Value}
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
