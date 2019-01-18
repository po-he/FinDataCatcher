package main

import(
	"fmt"

	"github.com/FinDataCatcher/base"
	"github.com/FinDataCatcher/data_catcher"
	"github.com/FinDataCatcher/tx_catcher"
	//"github.com/FinDataCatcher/utils"
)

type tx_catcher_quotes_data_job struct {}

func (j tx_catcher_quotes_data_job) RequestData(ctx *data_catcher.FinDataCatcherJobContext)  (string, error)	{
	bytes, err :=  base.DoHttpRequest(ctx.Url, 0)
	if nil != err {
		return "", err
	}

	return string(bytes), nil
	//return fmt.Sprintf("http://stock.gtimg.cn/data/index.php?appn=detail&action=data&c=%v&p=%v", 1, 1), nil
}
	
func (j tx_catcher_quotes_data_job) ProcessResponse(ctx *data_catcher.FinDataCatcherJobContext, resp string) bool {
	// YY-MM-DD 前缀
	ret, _, quotes := tx_catcher.TX_parse_stock_quotes_detail_response(resp)
	if false == ret {
		fmt.Printf("")
		return false
	}

	fmt.Printf("Receive quotes %v\n", quotes)

	// 在把自己放回去
	mng := data_catcher.GetGlobalFinDataCatcherMng()
	if nil == mng {
		return false
	}

	// 最后一次的数据还没有收到
	if quotes.Timestamp < ctx.MarketCloseTimestamp {
		mng.AddJobContext(ctx)
	}

	return true
}