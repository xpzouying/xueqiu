package xueqiu

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

// 7*24小时新闻
const liveNewsURL = "https://xueqiu.com/statuses/livenews/list.json?since_id=-1&max_id=-1&count=10"

type liveNews struct {
	ID        int    `json:"id"`
	Text      string `json:"text"`
	Target    string `json:"target"`
	CreatedAt int64  `json:"created_at"` // unix时间戳，单位ms
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

func GetLiveNews(ctx context.Context) (*RespLiveNews, error) {

	req, err := newGetLiveNewsRequest(liveNewsURL)
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
