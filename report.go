package xueqiu

import (
	"context"
	"time"

	"github.com/pkg/errors"
)

type CompanyReport struct {
	Title           string  `json:"title"`
	ReportCompany   string  `json:"rpt_comp"`
	RatingDesc      string  `json:"rating_desc"`
	TargetPriceMin  float32 `json:"target_price_min"`
	TargetPrinceMax float32 `json:"target_prince_max"`

	StatusID    int    `json:"status_id"`              // 猜测是这篇评论的id
	PubDate     int64  `json:"pub_date,omitempty"`     // 雪球正常返回结果，unix时间戳(ms)
	PublishDate string `json:"publish_date,omitempty"` // 使用日期返回
}

type RespCompanyReports struct {
	List []*CompanyReport `json:"list"`
}

const (
	reportURLPrefix = "https://stock.xueqiu.com/stock/report/latest.json?symbol="

	limitReportCount = 5
)

func (xq *Xueqiu) GetCompanyReport(ctx context.Context, stockSymbol string) ([]*CompanyReport, error) {
	if len(stockSymbol) == 0 {
		return nil, errors.New("empty stock code")
	}

	url := reportURLPrefix + stockSymbol

	var reports RespCompanyReports
	if err := xq.httpGetAndDecode(ctx, url, &reports); err != nil {
		return nil, err
	}

	validReports := reports.List
	if len(validReports) >= limitReportCount {
		validReports = reports.List[:limitReportCount]
	}

	// publish date不使用unix毫秒，使用正常日期
	for i := 0; i < len(validReports); i++ {
		unixSec := validReports[i].PubDate / 1000
		validReports[i].PublishDate = time.Unix(unixSec, 0).Format("2006-01-02")
	}

	return validReports, nil
}
