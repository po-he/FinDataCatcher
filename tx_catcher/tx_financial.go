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
	
	// 行情
	_TX_QUOTES_DATA_PREFIX = "v_"
	_MIN_TX_QUOTES_SPLITE_NUM  = 50  // 收盘后的长度
	_MAX_TX_QUOTES_SPLITE_NUM  = 54  // 实时的长度
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

	resp = strings.TrimRight(resp, ";")
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
行情數據請求url

http://qt.gtimg.cn/q=sh600519
*/
func TX_create_stock_quotes_detail_url(code string) string {
	return fmt.Sprintf("http://qt.gtimg.cn/q=%v", code)
}

/*
@param：
s ：YYYYMMDDHHMMSS  14位固定長度的時間字符串
    
@out:
對應的時間戳
*/
func tx_quotes_timestring_to_timestamp(s string) int64 {
	parser := func(str string) []string {
				l := make([]string, 0)
				if 14 != len(str) {
					return l
				}

				l = append(l, "")  // 這個只能預先填一個字段，應該是regex添加進的
				l = append(l, str[:4])
				l = append(l, str[4:6])
				l = append(l, str[6:8])
				l = append(l, str[8:10])
				l = append(l, str[10:12])
				l = append(l, str[12:])
				return l
			  }

	return utils.DateString2LocalNanoUnixTime(s, parser) / 1000000000
}

func TX_ShowQuotesTimestamp(s string) int64 {
	return tx_quotes_timestring_to_timestamp(s)
}


/*
@param：
resp: v_sh600519=
"1~贵州茅台~600519~358.74~361.29~361.88~27705~12252~15453~358.75~8~358.74~4~358.72~7~358.71~6~358.70~5~358.77~3~358.78~2~358.79~16~358.80~4~358.86~1~14:59:59/358.75/5/S/179381/28600|14:59:56/358.75/1/S/35875/28594|14:59:53/358.75/1/S/35875/28588|14:59:50/358.75/1/S/35875/28579|14:59:47/358.75/4/B/143499/28574|14:59:41/358.72/4/S/143501/28562~20170221150553~-2.55~-0.71~362.43~357.18~358.75/27705/994112865~27705~99411~0.22~27.24~~362.43~357.18~1.45~4506.49~4506.49~6.57~397.42~325.16~0.86";
分割符：~

格式：
 1  0: 未知
 2  1: 股票名字
 3  2: 股票代码
 4  3: 当前价格
 5  4: 昨收
 6  5: 今开
 7  6: 成交量（手）
 8  7: 外盘
 9  8: 内盘
10  9: 买一
11 10: 买一量（手）
12 11-18: 买二 买五
13 19: 卖一
14 20: 卖一量
15 21-28: 卖二 卖五
16 29: 最近逐笔成交
17 30: 时间
18 31: 涨跌
19 32: 涨跌%
20 33: 最高
21 34: 最低
22 35: 价格/成交量（手）/成交额
23 36: 成交量（手）
24 37: 成交额（万）
25 38: 换手率
26 39: 市盈率
27 40: 
28 41: 最高
29 42: 最低
30 43: 振幅
31 44: 流通市值
32 45: 总市值
33 46: 市净率
34 47: 涨停价
35 48: 跌停价
49-53 unknown

@out:
1-true/false
2-代码
3-行情数据

*/
func TX_parse_stock_quotes_detail_response(resp string) (bool, base.StockCode, base.StockQuotesDeatil) {
	code := base.StockCode{}
	quotes := base.StockQuotesDeatil{}
	resp = strings.TrimRight(resp, ";")
	v := strings.Split(resp, "=")
	if 2 != len(v) {
		return false, code, quotes
	}

	prefix := strings.Split(v[0], _TX_QUOTES_DATA_PREFIX)
	if 2 != len(prefix) {
		return false, code, quotes
	}

	// code
	code = tx_parse_stock_code(prefix[1])
	
	// paylod
	payload := strings.TrimPrefix(v[1], "\"")
	payload = strings.TrimRight(payload, "\"")
	payload_list := strings.Split(payload, "~")
	if len(payload_list) < _MIN_TX_QUOTES_SPLITE_NUM {
		return false, code, quotes
	}

	quotes.Price = utils.Stof64(payload_list[3])
	quotes.Close = utils.Stof64(payload_list[3])
	quotes.LastClose = utils.Stof64(payload_list[4])
	quotes.Open = utils.Stof64(payload_list[5])
	quotes.Amount = utils.StoI64(payload_list[6])
	quotes.Inside = utils.StoI64(payload_list[7])
	quotes.Outside = utils.StoI64(payload_list[8])

	// 9 ~ 29 忽略
	
	// 30 YYMMDDHHMMSS
	quotes.Timestamp = tx_quotes_timestring_to_timestamp(payload_list[30]) 
	//fmt.Printf("Payload date %v\n", payload_list[30])

	// 38: 换手率
	quotes.TR = utils.Stof64(payload_list[38]) // 換手率
	
	// 41 : 42
	quotes.Hight = utils.Stof64(payload_list[41])
	quotes.Low = utils.Stof64(payload_list[42])

	// 43: 振幅
	quotes.Amplitude = utils.Stof64(payload_list[43])

	// 47: 涨停价   48: 跌停价
	quotes.CeilPrice = utils.Stof64(payload_list[47]) 	
	quotes.FloorPrice = utils.Stof64(payload_list[48]) 	
	return true, code, quotes
}
