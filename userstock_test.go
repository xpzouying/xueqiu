package xueqiu

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserStocks(t *testing.T) {

	xq, err := NewWithEnvToken()
	assert.NoError(t, err)

	resp, err := xq.GetUserStocks(context.Background())
	assert.NoError(t, err)
	assert.NotNil(t, resp.Data)

	for _, stock := range resp.Data.Stocks {
		t.Logf("%+v", stock)
	}
}
