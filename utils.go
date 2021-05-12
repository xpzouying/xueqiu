package xueqiu

import "time"

type M map[string]interface{}

func makeTimestampMillisecond() int64 {
	return time.Now().Local().UnixNano() /
		(int64(time.Millisecond) / int64(time.Nanosecond))
}

// 传入的时间是否在设定的范围之类
func timeLessThan(timeMs int64, dur time.Duration) bool {

	t := time.Unix(0, timeMs*int64(time.Millisecond))
	now := time.Now()

	return now.Sub(t) <= dur
}
