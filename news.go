package xueqiu

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

const (
	// 7*24小时新闻
	liveNewsURL = "https://xueqiu.com/statuses/livenews/list.json?since_id=-1&max_id=-1&count=10"

	// 7*24重要新闻
	liveMarkNewsURL = "https://api.xueqiu.com/statuses/livenews/mark/list.json?max_id=-1&since_id=-1&size=20"
)

type liveNews struct {
	ID        int    `json:"id"`
	Text      string `json:"text"`
	Target    string `json:"target"`
	CreatedAt int64  `json:"created_at"` // unix时间戳，单位ms
	Mark      int    `json:"mark"`       // mark=1表示是重要新闻
}

type RespLiveNews struct {
	NextMaxID int         `json:"next_max_id,omitempty"` // 当前items中，最早的一条新闻的id。
	NextID    int         `json:"next_id,omitempty"`     // 目前还不知道是什么。
	Items     []*liveNews `json:"items,omitempty"`
}

func newGetLiveNewsRequest(url string) (req *http.Request, err error) {
	if req, err = http.NewRequest(http.MethodGet, url, nil); err != nil {
		return
	}
	req.AddCookie(xqCookie())
	return
}

// GetMarkLiveNews 获取7*24重要新闻
func GetMarkLiveNews(ctx context.Context) (*RespLiveNews, error) {

	return doGetLiveNews(ctx, liveMarkNewsURL)
}

// GetLiveNews 获取7*24新闻
func GetLiveNews(ctx context.Context) (*RespLiveNews, error) {

	return doGetLiveNews(ctx, liveNewsURL)
}

func doGetLiveNews(ctx context.Context, url string) (*RespLiveNews, error) {

	req, err := newGetLiveNewsRequest(url)
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if code := resp.StatusCode; code != http.StatusOK {
		return nil, errors.Errorf("get_live_news failed, status_code=%d", code)
	}

	var result RespLiveNews
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}
