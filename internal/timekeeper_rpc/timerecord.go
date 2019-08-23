package timekeeper_rpc

import (
	"log"
	"time"
)

func (r TimeRecord) GetDuration() time.Duration {
	timeStart := time.Unix(r.TimeStart, 0)
	timeEnd := time.Unix(r.TimeEnd, 0)
	log.Println("==>", timeStart, timeEnd)

	return timeEnd.Sub(timeStart)
}
