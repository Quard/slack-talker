package timekeeper_rpc

import (
	context "context"
	"errors"
	"io"
	"time"

	"github.com/getsentry/sentry-go"
	grpc "google.golang.org/grpc"
)

type TimeKeeper struct {
	uri string
}

func NewTimeKeeper(uri string) TimeKeeper {
	return TimeKeeper{uri: uri}
}

func (t TimeKeeper) Add(userID string, task string, timeStart time.Time, comment string) (*TimeRecord, error) {
	conn, errConn := t.getConnection()
	if errConn != nil {
		return nil, errors.New("connection between services error")
	}
	defer conn.Close()

	newRecord := &TimeStartRecord{
		UserID:    userID,
		Task:      task,
		TimeStart: timeStart.Unix(),
		Comment:   comment,
	}

	cl := NewTimeKeeperRPCClient(conn)
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	record, err := cl.Add(ctx, newRecord)
	if err != nil {
		sentry.CaptureException(err)
		return nil, errors.New("unable to add new record")
	}

	return record, nil
}

func (t TimeKeeper) Finish(userID string) (*TimeRecord, error) {
	conn, errConn := t.getConnection()
	if errConn != nil {
		return nil, errors.New("connection between services error")
	}
	defer conn.Close()

	cl := NewTimeKeeperRPCClient(conn)
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	record, err := cl.Done(ctx, &RecordID{UserID: userID})
	if err != nil {
		sentry.CaptureException(err)
		return nil, errors.New("unable to finish task")
	}

	return record, nil
}

func (t TimeKeeper) GetGroupedForDate(userID string, date time.Time) ([]*TimeGroupedRecord, error) {
	conn, errConn := t.getConnection()
	if errConn != nil {
		return nil, errors.New("connection between services error")
	}
	defer conn.Close()

	cl := NewTimeKeeperRPCClient(conn)
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	stream, err := cl.GetGroupedForDate(ctx, &GetForDate{UserID: userID, Date: date.Unix()})
	if err != nil {
		sentry.CaptureException(err)
		return nil, err
	}

	report := []*TimeGroupedRecord{}
	for {
		record, err := stream.Recv()
		if err != nil {
			if err != io.EOF {
				sentry.CaptureException(err)
			}

			break
		}
		report = append(report, record)
	}

	return report, nil
}

func (t TimeKeeper) getConnection() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(t.uri, grpc.WithInsecure())
	if err != nil {
		sentry.CaptureException(err)
		return nil, err
	}

	return conn, nil
}
