package xueqiu

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimeLessThan(t *testing.T) {

	ts := []struct {
		testname string

		timeMs int64
		dur    time.Duration

		want bool
	}{
		{
			testname: "1小时前，是否在2小时 以内",
			timeMs:   (time.Now().Add(-1 * time.Hour).Unix()) * 1000,
			dur:      2 * time.Hour,

			want: true,
		},
		{
			testname: "3小时前，是否在2小时 以内",
			timeMs:   (time.Now().Add(-3 * time.Hour).Unix()) * 1000,
			dur:      2 * time.Hour,

			want: false,
		},
	}

	for _, tc := range ts {
		t.Run(tc.testname, func(t *testing.T) {
			got := timeLessThan(tc.timeMs, tc.dur)

			assert.Equal(t, tc.want, got)
		})
	}
}
