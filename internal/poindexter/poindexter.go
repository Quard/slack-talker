package poindexter

import (
	context "context"
	"errors"
	"io"
	"time"

	"github.com/getsentry/sentry-go"
	grpc "google.golang.org/grpc"

	"github.com/Quard/slack-talker/internal/authority"
)

type Poindexter struct {
	uri string
}

func NewPoindexter(uri string) Poindexter {
	return Poindexter{uri: uri}
}

func (p Poindexter) AddReadingListRecord(user *authority.User, url string) (*ReadingListRecord, error) {
	conn, errConn := p.getConnection()
	if errConn != nil {
		return nil, errors.New("connection between services error")
	}
	defer conn.Close()

	cl := NewInternalAPIClient(conn)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	record, err := cl.Add(ctx, &URL{UserID: user.ID, Url: url})
	if err != nil {
		sentry.CaptureException(err)
		return nil, errors.New("unable to store url")
	}

	return record, nil
}

func (p Poindexter) MarkRecordAsRead(user *authority.User, objID string) (*ReadingListRecord, error) {
	conn, errConn := p.getConnection()
	if errConn != nil {
		return nil, errors.New("connection between services error")
	}
	defer conn.Close()

	cl := NewInternalAPIClient(conn)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	record, err := cl.MarkAsRead(ctx, &ID{UserID: user.ID, Id: objID})
	if err != nil {
		sentry.CaptureException(err)
		return nil, errors.New("unable to update record")
	}

	return record, nil
}

func (p Poindexter) ListReadingListRecord(user *authority.User, ch chan *ReadingListRecord) error {
	conn, errConn := p.getConnection()
	if errConn != nil {
		return errors.New("connection between services error")
	}

	cl := NewInternalAPIClient(conn)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	stream, err := cl.List(ctx, &User{ID: user.ID})
	if err != nil {
		sentry.CaptureException(err)
		return err
	}

	go func() {
		for {
			record, err := stream.Recv()
			if err != nil {
				if err != io.EOF {
					sentry.CaptureException(err)
				}
				close(ch)
				defer conn.Close()

				return
			}

			ch <- record
		}
	}()

	return nil
}

func (p Poindexter) getConnection() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(p.uri, grpc.WithInsecure())
	if err != nil {
		sentry.CaptureException(err)
		return nil, err
	}

	return conn, nil
}
