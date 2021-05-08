package xueqiu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {

	_, err := NewWithEnvToken()

	assert.NoError(t, err)
}
