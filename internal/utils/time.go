package utils

import "time"

func NowTs() int64 {
	return time.Now().UnixMilli()
}
