package xueqiu

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

// 7*24小时新闻
const liveNewsURL = "https://xueqiu.com/statuses/livenews/list.json?since_id=-1&max_id=-1&count=10"

func newGetLiveNewsRequest(url string) (req *http.Request, err error) {
	if req, err = http.NewRequest(http.MethodGet, url, nil); err != nil {
		return
	}
	req.AddCookie(xqCookie())
	return
}

func GetLiveNews(ctx context.Context) (M, error) {

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

	var m M
	if err := json.NewDecoder(resp.Body).Decode(&m); err != nil {
		return nil, err
	}

	return m, nil
}
