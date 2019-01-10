package base

import(
	"fmt"
)

const (
	DEAL_SIDE_BUY  = 1
	DEAL_SIDE_SELL = 2	

	STOCK_MARKET_ID_SZ      = 1
	STOCK_MARKET_STR_SZ     = "SZ"

	STOCK_MARKET_ID_SH      = 2
	STOCK_MARKET_STR_SH     = "SH"

	STOCK_BLOCK_ID_SZ_A    = 0
	STOCK_BLOCK_ID_SH_A    = 6
	STOCK_BLOCK_ID_CY      = 3   

	STOCK_BLOCK_CY_BASE_ID = 300000
)

func IsMarketSZ(c *StockCode) bool {
	if nil != c && c.Market == STOCK_MARKET_ID_SZ {
		return true
	} else {
		return false
	}
}

func IsMarketSH(c *StockCode) bool {
	if nil != c && c.Market == STOCK_MARKET_ID_SH {
		return true
	} else {
		return false
	}	
}

func IsMarketBlockSH(c *StockCode) bool {
	return IsMarketSH(c)
}

func IsMarckBlockSZ(c *StockCode) bool {
	if nil != c && c.Market == STOCK_MARKET_ID_SZ && c.Code < STOCK_BLOCK_CY_BASE_ID {
		return true
	} else {
		return false
	}
}

func IsMarketBlockCY(c *StockCode) bool {
	if nil != c && c.Market == STOCK_MARKET_ID_SZ && c.Code >= STOCK_BLOCK_CY_BASE_ID {
		return true
	} else {
		return false
	}
}


type StockCode struct {
	Market int32      "json:Market"
	Code   int32      "json:Code"
}

func (c StockCode) String() string {
	if STOCK_MARKET_ID_SZ == c.Market {
		return fmt.Sprintf("%v %06d", STOCK_MARKET_STR_SZ, c.Code)
	} else if STOCK_MARKET_ID_SH == c.Market {
		return fmt.Sprintf("%v %06d", STOCK_MARKET_STR_SH, c.Code)
	} else {
		return ""
	}
}

func (c StockCode) Detail() string {
	if STOCK_MARKET_ID_SZ == c.Market {
		if c.Code < STOCK_BLOCK_CY_BASE_ID {
			return fmt.Sprintf("%v %06d %v", STOCK_MARKET_STR_SZ, c.Code, STOCK_BLOCK_ID_SZ_A)
		} else {
			return fmt.Sprintf("%v %06d %v", STOCK_MARKET_STR_SZ, c.Code, STOCK_BLOCK_ID_CY)
		}
	} else if STOCK_MARKET_ID_SH == c.Market {
		return fmt.Sprintf("%v %06d %v", STOCK_MARKET_STR_SH, c.Code, STOCK_BLOCK_ID_SH_A)
	} else {
		return ""
	}
}

type StockTickDealDetail struct {
	Timestamp int64   "json:Timestamp"
	Time      string  "json:Time"
	Price     float64 "json:Price"
	Amount    int64   "json:Amount"
	Total     int64   "json:Total"
	Side      int8    "json:Side"
}


