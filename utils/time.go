package utils

import "time"

const timeLayout = "2006-01-02 15:04:05"

func Unix2Time(unix int64) time.Time {
	return time.Unix(unix, 0).UTC()
}
