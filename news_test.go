package xueqiu

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLiveNews(t *testing.T) {
	xq, err := NewWithEnvToken()
	assert.NoError(t, err)

	ts := []struct {
		testname string

		fn func(ctx context.Context) (*RespLiveNews, error)
	}{
		{"获取7*24小时的新闻", xq.GetLiveNews},
		{"获取7*24小时的重要新闻", xq.GetMarkLiveNews},
	}

	for _, tc := range ts {
		t.Run(tc.testname, func(t *testing.T) {
			res, err := tc.fn(context.Background())

			assert.NoError(t, err)
			assert.NotNil(t, res)
			assert.NotNil(t, res.Items)

		})
	}
}
