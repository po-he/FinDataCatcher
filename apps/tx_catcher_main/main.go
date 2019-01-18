package main

import(
	"fmt"

	"github.com/FinDataCatcher/base"
	"github.com/FinDataCatcher/data_catcher"
	"github.com/FinDataCatcher/tx_catcher"
	"github.com/FinDataCatcher/utils"
)

var (
	//_D base.FinDataCatcherMng
	_quit_flag = false
)


func init() {
	//_D = new(base.FinDataCatcherMng)
}

func main() {
	// 1. 
	mng := data_catcher.GetGlobalFinDataCatcherMng()
	if nil == mng {
		fmt.Printf("Invalid GlobalFinDataCatcherMng\n")
		return
	}

	mng.Init(2)

	quote_job_sh600519 := data_catcher.FinDataCatcherJobContext {
										DownloadInfo : new(base.DailyDownloadeInfo),
										Job : tx_catcher_quotes_data_job{},
										Url : tx_catcher.TX_create_stock_quotes_detail_url("sh600519"),
										MarketCloseTimestamp : utils.NormalDate2LocalUnixTimeSec("2019-01-17 15:00:00"),
				   		  			}

	quote_job_sh600520 := data_catcher.FinDataCatcherJobContext {
										DownloadInfo : new(base.DailyDownloadeInfo),
										Job : tx_catcher_quotes_data_job{},
										Url : tx_catcher.TX_create_stock_quotes_detail_url("sh600520"),
										MarketCloseTimestamp : utils.NormalDate2LocalUnixTimeSec("2019-01-17 15:00:00"),
				   		  			}

	quote_job_sh600521 := data_catcher.FinDataCatcherJobContext {
										DownloadInfo : new(base.DailyDownloadeInfo),
										Job : tx_catcher_quotes_data_job{},
										Url : tx_catcher.TX_create_stock_quotes_detail_url("sh600521"),
										MarketCloseTimestamp : utils.NormalDate2LocalUnixTimeSec("2019-01-17 15:00:00"),
				   		  
									}			
	quote_job_sh600524 := data_catcher.FinDataCatcherJobContext {
										DownloadInfo : new(base.DailyDownloadeInfo),
										Job : tx_catcher_quotes_data_job{},
										Url : tx_catcher.TX_create_stock_quotes_detail_url("sh600524"),
										MarketCloseTimestamp : utils.NormalDate2LocalUnixTimeSec("2019-01-17 15:00:00"),
									}
						
	mng.AddJobContext(&quote_job_sh600519)									
	mng.AddJobContext(&quote_job_sh600520)									
	mng.AddJobContext(&quote_job_sh600521)									
	mng.AddJobContext(&quote_job_sh600524)									


	loop()
	
	// 2. 获取

	// 3. 初始化CatcherManager, 创建对应数量的CatcherRoutine

	// 4. 更新数据列表

	// 5. 读取

	// 
}

func loop() {
	mng := data_catcher.GetGlobalFinDataCatcherMng()
	if nil == mng {
		return
	}
		
	for {
		if _quit_flag == true  {
			break
		}

		job_num := mng.GetJobContextNum()
		catcher_num := mng.GetHttpCatcherNum()
		if 0 == job_num || 0 == catcher_num {
			continue
		}

		// 这里需要定时来走


		// 获取
		ctx_list := mng.PopJobContext(catcher_num)
		for _, ctx := range ctx_list {
			if false == mng.Dispatch(ctx) {
				fmt.Printf("Diapatch job ctx failed! need requeue ctx\n")
				continue
			}
		}
	}
}