package xueqiu

import (
	"context"
)

const (
	// 7*24小时新闻
	liveNewsURL = "https://xueqiu.com/statuses/livenews/list.json?since_id=-1&max_id=-1&count=10"

	// 7*24重要新闻
	liveMarkNewsURL = "https://api.xueqiu.com/statuses/livenews/mark/list.json?max_id=-1&since_id=-1&size=20"
)

type LiveNews struct {
	ID        int    `json:"id"`
	Text      string `json:"text"`
	Target    string `json:"target"`
	CreatedAt int64  `json:"created_at"` // unix时间戳，单位ms
	Mark      int    `json:"mark"`       // mark=1表示是重要新闻
}

type RespLiveNews struct {
	NextMaxID int         `json:"next_max_id,omitempty"` // 当前items中，最早的一条新闻的id。
	NextID    int         `json:"next_id,omitempty"`     // 目前还不知道是什么。
	Items     []*LiveNews `json:"items,omitempty"`
}

// GetMarkLiveNews 获取7*24重要新闻
func (xq *Xueqiu) GetMarkLiveNews(ctx context.Context) (*RespLiveNews, error) {

	return xq.doGetLiveNews(ctx, liveMarkNewsURL)
}

// GetLiveNews 获取7*24新闻
func (xq *Xueqiu) GetLiveNews(ctx context.Context) (*RespLiveNews, error) {

	return xq.doGetLiveNews(ctx, liveNewsURL)
}

func (xq *Xueqiu) doGetLiveNews(ctx context.Context, url string) (*RespLiveNews, error) {
	var result RespLiveNews
	if err := xq.httpGetAndDecode(ctx, url, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
