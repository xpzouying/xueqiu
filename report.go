package xueqiu

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

type CompanyReport struct {
	Title           string  `json:"title"`
	ReportCompany   string  `json:"rpt_comp"`
	RatingDesc      string  `json:"rating_desc"`
	TargetPriceMin  float32 `json:"target_price_min"`
	TargetPrinceMax float32 `json:"target_prince_max"`

	PubDate     int64  `json:"pub_date,omitempty"`     // 雪球正常返回结果，unix时间戳(ms)
	PublishDate string `json:"publish_date,omitempty"` // 使用日期返回
}

type RespCompanyReports struct {
	List []*CompanyReport `json:"list"`
}

func newGetRequest(url string) (req *http.Request, err error) {
	if req, err = http.NewRequest(http.MethodGet, url, nil); err != nil {
		return
	}
	req.AddCookie(xqCookie())

	return
}

const reportURLPrefix = "https://stock.xueqiu.com/stock/report/latest.json?symbol="

func GetCompanyReport(ctx context.Context, stock string) ([]*CompanyReport, error) {

	url := reportURLPrefix + stock

	req, err := newGetRequest(url)
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if code := resp.StatusCode; code != http.StatusOK {
		return nil, errors.Errorf("http status code not ok: %v", code)
	}

	var reports RespCompanyReports
	if err := json.NewDecoder(resp.Body).Decode(&reports); err != nil {
		return nil, err
	}

	// publish date不使用unix毫秒，使用正常日期
	list := reports.List
	for i := 0; i < len(list); i++ {
		unixSec := list[i].PubDate / 1000
		list[i].PublishDate = time.Unix(unixSec, 0).Format("2006-01-02")
		list[i].PubDate = 0
	}

	return list, nil
}
