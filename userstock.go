package xueqiu

import "context"

const (
	// 自选股票
	userFavoriteStockURL = "https://stock.xueqiu.com/v5/stock/portfolio/stock/list.json?size=1000&category=1&pid=-1"
)

type favStock struct {
	Stock

	Remark   string `json:"remark"`
	Exchange string `json:"exchange"`
	Created  int64  `json:"created"`
}

type RespUserStocks struct {
	ErrorCode        int    `json:"error_code"`
	ErrorDescription string `json:"error_description"`

	Data *struct {
		Pid      int `json:"pid"`
		Category int `json:"category"`

		Stocks []*favStock `json:"stocks"`
	} `json:"data"`
}

func (xq *Xueqiu) GetUserStocks(ctx context.Context) (*RespUserStocks, error) {

	var resp RespUserStocks
	if err := xq.httpGetAndDecode(ctx, userFavoriteStockURL, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
