package xueqiu

import "time"

type M map[string]interface{}

func makeTimestampMillisecond() int64 {
	return time.Now().Local().UnixNano() /
		(int64(time.Millisecond) / int64(time.Nanosecond))
}
