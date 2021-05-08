package xueqiu

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/pkg/errors"
)

// TokenEnvKey 环境变量中的token字段名字
const TokenEnvKey = "XQTOKEN"

type Xueqiu struct {
	token string
	c     *http.Client
}

// New to create a xueqiu robot. To follow the step to get a token: https://blog.crackcreed.com/diy-xue-qiu-app-shu-ju-api/
func New(token string) (*Xueqiu, error) {

	xq := &Xueqiu{
		token: token,
		c:     &http.Client{},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := xq.CheckHealth(ctx); err != nil {
		return nil, err
	}

	return xq, nil
}

// NewWithEnvToken must get token from os.Env("XQTOKEN").
func NewWithEnvToken() (*Xueqiu, error) {
	token := os.Getenv(TokenEnvKey)
	if len(token) == 0 {
		return nil,
			errors.Errorf("get os env error: %s", TokenEnvKey)
	}

	return New(token)
}

func (xq *Xueqiu) httpCookie() *http.Cookie {
	return &http.Cookie{
		Name:  "xq_a_token",
		Value: xq.token,
	}
}

func (xq *Xueqiu) httpGet(ctx context.Context, url string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.AddCookie(xq.httpCookie())

	req = req.WithContext(ctx)

	return xq.c.Do(req)
}

func (xq *Xueqiu) httpGetAndDecode(ctx context.Context, url string, v interface{}) error {
	resp, err := xq.httpGet(ctx, url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if code := resp.StatusCode; code != http.StatusOK {
		return errors.Errorf("http get return invalid code: %v", code)
	}

	return json.NewDecoder(resp.Body).Decode(&v)
}

func (xq *Xueqiu) CheckHealth(ctx context.Context) error {
	// 使用7*24小时新闻作为token合法性的校验
	url := "https://xueqiu.com/statuses/livenews/list.json?since_id=-1&max_id=-1&count=1"

	resp, err := xq.httpGet(ctx, url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if code := resp.StatusCode; code != http.StatusOK {
		return errors.Errorf("check health error: status_code=%d", code)
	}

	return nil
}
