package tx_catcher

import(
	"strings"
	"fmt"

	"github.com/FinDataCatcher/utils"
	"github.com/FinDataCatcher/base"
)

//http://stock.gtimg.cn/data/index.php?appn=detail&action=data&c=sh600550

const (
	_TX_TICK_DETAIL_PREFIX = "v_detail_data_"
	_TX_STATISTIC_DATA_PREFIX = ""
)

/*
@param：
s: 	70/09:33:27/12.78/-0.01/251/321027/S 
格式：分时索引/时间/成交价格/价格变动/成交量(手)/成交价格/成交方向
*/
func tx_parse_stock_tick_detail_string(s string) (bool, base.StockTickDealDetail) {
	v := strings.Split(s, "/")	
	if len(v) != 7 {
		return false, base.StockTickDealDetail{}
	}

	detail := base.StockTickDealDetail{
					Time : v[1],
					Price : utils.Stof64(v[2]),
					Amount : utils.StoI64(v[4]),
					Total : utils.StoI64(v[5]),
	}

	side := v[6]
	if side == "S" {
		detail.Side = base.DEAL_SIDE_SELL
	} else if side == "B" {
		detail.Side = base.DEAL_SIDE_BUY
	}

	return true, detail
}

/*
@param：
sz002451
sh

@out:
*/
func tx_parse_stock_code(s string) base.StockCode {
	index := strings.Index(s, "sz")
	if -1 != index {
		//fmt.Printf("sz index %v last %v\n", index, s[index+2:])
		return base.StockCode {
					Market : base.STOCK_MARKET_ID_SZ,
					Code   : utils.StoI32(s[index+2:]),
			   }
	}
	
	index = strings.Index(s, "sh")
	if -1 != index {
		//fmt.Printf("sh index %v last %v\n", index, s[index+2:])
		return base.StockCode {
					Market : base.STOCK_MARKET_ID_SH,
					Code   : utils.StoI32(s[index+2:]),
			   }	
	}

	//fmt.Printf("sh index %v\n", index)
	return base.StockCode{}
}

/*
@param：
code ：交易所簡碼+6位編號
index ：索引
    
@out:
tick數據請求url
*/
func TX_create_stock_tick_detail_url(code string, index int) string {
	return fmt.Sprintf("http://stock.gtimg.cn/data/index.php?appn=detail&action=data&c=%v&p=%v", code, index)
}

/*
@param：
resp: 	v_detail_data_sz002451=[1,"70/09:33:27/12.78/-0.01/251/321027/S|71/09:33:31/12.80/0.02/392/501495/B"]
格式：前缀=[索引, "分时成交信息|分时成交信息"]

dateYMD: "2017-12-01"
    
@out:
1-true/false
2-代码
3-索引
4-分时数据列表
*/
func TX_parse_stock_tick_detail_response(resp string, dateYMD string) (bool, base.StockCode, int32, []base.StockTickDealDetail) {
	index := int32(-1)
	code := base.StockCode{}
	detail := make([]base.StockTickDealDetail, 0)

	v := strings.Split(resp, "=")
	if len(v) != 2 {
		return false, code, index, detail
	}

	prefix := strings.Split(v[0], _TX_TICK_DETAIL_PREFIX)
	if len(prefix) != 2 {
		return false, code, index, detail
	}

	//code_str := prefix[1]
	code = tx_parse_stock_code(prefix[1])

	payload := strings.TrimPrefix(v[1], "[")
	payload = strings.TrimRight(payload, "]")

	content := strings.Split(payload, ",")
	if len(content) != 2 {
		return false, code, index, detail	
	}

	index = utils.StoI32(content[0])
	detail_str := strings.TrimPrefix(content[1], "\"")
	detail_str = strings.TrimRight(detail_str, "\"")

	detail_list := strings.Split(detail_str, "|") 
	for _, s := range detail_list {
		res, d := tx_parse_stock_tick_detail_string(s)
		if true == res {
			d.Time = fmt.Sprintf("%s %s", dateYMD, d.Time)
			d.Timestamp = utils.NormalDate2LocalUnixTimeSec(d.Time)

			detail = append(detail, d)
		}
	}
	
	return true, code, index, detail
}

/*
@param：
code ：交易所簡碼+6位編號
    
@out:
統計數據請求url
*/
func TX_create_stock_statistic_data_url(code string) string {
	//return fmt.Sprintf("http://stock.gtimg.cn/data/index.php?appn=detail&action=data&c=%v&p=%v", code, index)
	return ""
}

/*
@param：
resp: 	v_detail_data_sz002451=[1,"70/09:33:27/12.78/-0.01/251/321027/S|71/09:33:31/12.80/0.02/392/501495/B"]
格式：前缀=[索引, "分时成交信息|分时成交信息"]

dateYMD: "2017-12-01"
    
@out:
1-true/false
2-代码
3-索引
4-分时数据列表
*/
func TX_parse_stock_statistic_data_response(resp string, dateYMD string) (bool, base.StockCode) {
	return false, base.StockCode{}
}
