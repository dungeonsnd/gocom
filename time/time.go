package time

import "time"

func TimestampMSec() int64 {
	now := time.Now()
	timestamp := now.UnixNano() / int64(time.Millisecond)
	return timestamp
}
