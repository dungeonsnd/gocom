package time

import "time"

func TimestampMSec() int64 {
	now := time.Now()
	timestamp := now.UnixNano() / int64(time.Millisecond)
	return timestamp
}

func CheckTimestamp(timestamp int64, minMSec int64) bool {
	if TimestampMSec()-timestamp > minMSec {
		return false
	} else {
		return true
	}
}
