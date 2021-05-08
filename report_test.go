package xueqiu

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCompanyReport(t *testing.T) {

	xq, err := NewWithEnvToken()
	assert.NoError(t, err)

	reports, err := xq.GetCompanyReport(context.Background(), "SZ300750") // 宁德时代
	assert.NoError(t, err)
	assert.NotNil(t, reports)

	for _, report := range reports {
		t.Logf("report: %+v", report)
	}
}
