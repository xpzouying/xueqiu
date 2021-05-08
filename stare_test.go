package xueqiu

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeGetStareURL(t *testing.T) {

	cateTypes := []CategoryType{CateDynamic, CateEvent}

	for _, cateType := range cateTypes {
		url, err := makeGetStareURL(cateType)

		assert.NoError(t, err)
		assert.Contains(t, url, cateType)
	}
}

func TestGetStareItems(t *testing.T) {
	ts := []struct {
		testname string

		fn func(ctx context.Context) (*RespStareItem, error)
	}{
		{"获取关注的异常波动", GetDynamicStareItems},
		{"获取关注的重大事件", GetEventStareItems},
	}

	for _, tc := range ts {

		t.Run(tc.testname, func(t *testing.T) {
			items, err := tc.fn(context.Background())

			assert.NoError(t, err)
			assert.NotNil(t, items)
		})
	}
}
