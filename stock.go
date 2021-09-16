package xueqiu

// StockType 股票的上市地区
type StockType uint16

const (
	StockTypeUS = 0
	StockTypeCN = 11
	StockTypeHK = 30
)

type Stock struct {
	Symbol string    `json:"symbol"` // 股票代码（SZ000858，SH600519)
	Name   string    `json:"name"`   // 股票中文名字
	Type   StockType `json:"type"`   // 股票上市地方
}
