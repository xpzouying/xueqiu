package xueqiu

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

const countOfStare = 10 // 获取的item个数限制，app中默认为50。

type CategoryType string

const (
	CateDynamic CategoryType = "dynamic"
	CateEvent   CategoryType = "event"
)

// StareItemType 股票的上市地区
type StareItemType uint16

const (
	StareItemUS = 0
	StareItemSZ = 11
	StareItemHK = 30
)

type StareItem struct {
	Timestamp    int64         `json:"timestamp"`  // 时间戳。单位ms
	Symbol       string        `json:"symbol"`     // 股票代码
	Name         string        `json:"name"`       // 股票中文名字
	StareKey     string        `json:"stare_key"`  // 小icon，fall。
	StareType    string        `json:"stare_type"` // 小icon，跌幅。
	Text         string        `json:"text"`
	Category     CategoryType  `json:"category"` // dynamic or event
	Type         StareItemType `json:"type"`
	CurrentPrice float32       `json:"current"` // 当前的价钱
	Percent      float32       `json:"percent"` // 涨跌幅
}

type RespStareItem struct {
	Data struct {
		Items []*StareItem `json:"items"`
	} `json:"data"`
}

// GetDynamicStareItem 获取关注动态
func GetDynamicStareItems(ctx context.Context) (*RespStareItem, error) {

	return getStareItems(ctx, CateDynamic)
}

// GetEventStareItem 获取关注动态
func GetEventStareItems(ctx context.Context) (*RespStareItem, error) {

	return getStareItems(ctx, CateEvent)
}

func getStareItems(ctx context.Context, cateType CategoryType) (*RespStareItem, error) {
	// url
	// https://stock.xueqiu.com/v5/stock/stare/list.json?
	//	count=-50&start=1620397068604&type=event

	u, err := makeGetStareURL(cateType)
	if err != nil {
		return nil, err
	}

	req, err := newGetRequestWithCookie(u)
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var res RespStareItem
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
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

	q.Set("start", fmt.Sprintf("%v", time.Now().Local().UnixNano()))
	q.Set("type", string(cateType))

	u.RawQuery = q.Encode()

	return u.String(), nil
}
