package xueqiu

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLiveNews(t *testing.T) {

	m, err := GetLiveNews(context.Background())

	assert.NoError(t, err)
	assert.NotNil(t, m)
	assert.NotNil(t, m["items"])
}
