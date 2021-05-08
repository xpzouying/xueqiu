package xueqiu

import (
	"net/http"
	"os"
	"strings"
)

var (
	c *http.Client

	xqtoken string
)

func init() {
	token := strings.TrimSpace(os.Getenv("XQTOKEN"))
	if len(token) == 0 {
		panic("must set env: XQTOKEN")
	}
	xqtoken = token

	c = &http.Client{}
}

func xqCookie() *http.Cookie {
	return &http.Cookie{
		Name:  "xq_a_token",
		Value: xqtoken,
	}
}

func newGetRequestWithCookie(url string) (req *http.Request, err error) {
	if req, err = http.NewRequest(http.MethodGet, url, nil); err != nil {
		return
	}
	req.AddCookie(xqCookie())

	return
}
