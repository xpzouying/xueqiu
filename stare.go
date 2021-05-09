package xueqiu

import (
	"context"
	"fmt"
	"net/url"
)

const countOfStare = 10 // 获取的item个数限制，app中默认为50。

type CategoryType string

const (
	CateDynamic CategoryType = "dynamic"
	CateEvent   CategoryType = "event"
)

type StareItem struct {
	Stock

	// Symbol       string        `json:"symbol"`     // 股票代码
	// Name         string        `json:"name"`       // 股票中文名字
	// Type         StareItemType `json:"type"`

	Timestamp    int64        `json:"timestamp"`  // 时间戳。单位ms
	StareKey     string       `json:"stare_key"`  // 小icon，fall。
	StareType    string       `json:"stare_type"` // 小icon，跌幅。
	Text         string       `json:"text"`
	Category     CategoryType `json:"category"` // dynamic or event
	CurrentPrice float32      `json:"current"`  // 当前的价钱
	Percent      float32      `json:"percent"`  // 涨跌幅
}

type RespStareItem struct {
	Data struct {
		Items []*StareItem `json:"items"`
	} `json:"data"`
}

// GetDynamicStareItem 获取关注动态
func (xq *Xueqiu) GetDynamicStareItems(ctx context.Context) (*RespStareItem, error) {

	return xq.getStareItems(ctx, CateDynamic)
}

// GetEventStareItem 获取关注动态
func (xq *Xueqiu) GetEventStareItems(ctx context.Context) (*RespStareItem, error) {

	return xq.getStareItems(ctx, CateEvent)
}

func (xq *Xueqiu) getStareItems(ctx context.Context, cateType CategoryType) (*RespStareItem, error) {
	// url
	// https://stock.xueqiu.com/v5/stock/stare/list.json?
	//	count=-50&start=1620397068604&type=event

	url, err := makeGetStareURL(cateType)
	if err != nil {
		return nil, err
	}

	var res RespStareItem
	if err := xq.httpGetAndDecode(ctx, url, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func makeGetStareURL(cateType CategoryType) (string, error) {

	u, err := url.Parse(
		fmt.Sprintf("https://stock.xueqiu.com/v5/stock/stare/list.json?count=-%d",
			countOfStare,
		),
	)
	if err != nil {
		return "", err
	}

	q := u.Query()

	q.Set("start", fmt.Sprintf("%v", makeTimestampMillisecond()))
	q.Set("type", string(cateType))

	u.RawQuery = q.Encode()

	return u.String(), nil
}
