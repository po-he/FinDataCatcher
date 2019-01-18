package main

import (
	"fmt"

	"github.com/FinDataCatcher/base"
	"github.com/FinDataCatcher/tx_catcher"
	"github.com/FinDataCatcher/utils"
	"github.com/FinDataCatcher/data_catcher"
	
	//"github.com/syndtr/goleveldb/leveldb"
)

type tx_catcher_tick_data_job struct {}


func (j tx_catcher_tick_data_job) RequestData(ctx *data_catcher.FinDataCatcherJobContext)  (string, error)	{
	//return fmt.Sprintf("http://stock.gtimg.cn/data/index.php?appn=detail&action=data&c=%v&p=%v", 1), bkl
	bytes, err :=  base.DoHttpRequest(ctx.Url, 0)
	if nil != err {
		return "", err
	}

	return string(bytes), nil
}
	
func (j tx_catcher_tick_data_job) ProcessResponse(ctx *data_catcher.FinDataCatcherJobContext, resp string) bool {
	//ctx
	//info *DailyDownloadeInfo
	download_info := ctx.DownloadInfo

	// YY-MM-DD 前缀
	date_str := utils.CurrentLocalDate()
	ret, _, index, tick_list := tx_catcher.TX_parse_stock_tick_detail_response(resp, date_str)
	if false == ret {
		fmt.Printf("")
		return false
	}

	// Tick數據保存到DB
	for _, tick_data := range tick_list {
		fmt.Printf("%v\n", tick_data)
	}

	// 記錄下一次索引
	download_info.LastTickDataIndex = index
	
	// 保存一次
	ctx.SaveDailyDownloadInfo()
	return true
}
