package xueqiu

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserStocks(t *testing.T) {

	xq, err := NewWithEnvToken()
	assert.NoError(t, err)

	stocks, err := xq.GetUserStocks(context.Background())
	assert.NoError(t, err)
	assert.NotNil(t, stocks)

	t.Logf("stock cate: category=%v pid=%v", stocks.Category, stocks.Pid)
	for _, stock := range stocks.Stocks {
		t.Logf("%+v", stock)
	}
}

func TestGetUserFollowReports(t *testing.T) {
	xq, err := NewWithEnvToken()
	assert.NoError(t, err)

	ctx := context.Background()

	reports, err := xq.GetUserFollowReports(ctx)
	assert.NoError(t, err)

	for symbol, report := range reports {

		t.Logf("synbol:%v, %+v, reports:\n", symbol, report.FavStock)

		for _, r := range report.CompanyReports {
			t.Logf("company: %+v", r)
		}
	}
}
