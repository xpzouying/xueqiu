package xueqiu

import (
	"context"

	"github.com/pkg/errors"
)

const (
	// 自选股票
	userFavoriteStockURL = "https://stock.xueqiu.com/v5/stock/portfolio/stock/list.json?size=1000&category=1&pid=-1"
)

type FavStock struct {
	Stock

	Remark   string `json:"remark"`
	Exchange string `json:"exchange"`
	Created  int64  `json:"created"`
}

type UserStocks struct {
	Pid      int `json:"pid"`
	Category int `json:"category"`

	Stocks []*FavStock `json:"stocks"`
}

type RespUserStocks struct {
	ErrorCode        int    `json:"error_code"`
	ErrorDescription string `json:"error_description"`

	Item *UserStocks `json:"data"`
}

func (xq *Xueqiu) GetUserStocks(ctx context.Context) (*UserStocks, error) {

	var resp RespUserStocks
	if err := xq.httpGetAndDecode(ctx, userFavoriteStockURL, &resp); err != nil {
		return nil, err
	}

	if resp.ErrorCode != 0 {
		return nil, errors.New(resp.ErrorDescription)
	}

	return resp.Item, nil
}

type StockReports struct {
	FavStock       *FavStock
	CompanyReports []*CompanyReport
}

// GetUserFollowReports 获取用户关注的公司报告。
// 返回: <stock_symbol: reports>
func (xq *Xueqiu) GetUserFollowReports(ctx context.Context) (map[string]StockReports, error) {

	userStocks, err := xq.GetUserStocks(ctx)
	if err != nil {
		return nil, err
	}

	stocks := userStocks.Stocks
	result := make(map[string]StockReports, len(stocks))

	for _, stock := range userStocks.Stocks {
		reports, err := xq.GetCompanyReport(ctx, stock.Symbol)
		if err != nil {
			return nil, err
		}

		result[stock.Symbol] = StockReports{
			FavStock:       stock,
			CompanyReports: reports,
		}
	}

	return result, nil
}
