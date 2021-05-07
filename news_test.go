package xueqiu

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLiveNews(t *testing.T) {

	res, err := GetLiveNews(context.Background())

	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.NotNil(t, res.Items)
}
