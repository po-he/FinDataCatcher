package main

import(
	"github.com/FinDataCatcher/base"
	//"github.com/FinDataCatcher/tx_catcher"
	//"github.com/FinDataCatcher/utils"
)

type tx_catcher_quotes_data_job struct {}


func (j tx_catcher_quotes_data_job) RequestData(ctx *base.FinDataCatcherJobContext)  (string, error)	{
	bytes, err :=  base.DoHttpRequest(ctx.Url, 0)
	if nil != err {
		return "", err
	}

	return string(bytes), nil
	//return fmt.Sprintf("http://stock.gtimg.cn/data/index.php?appn=detail&action=data&c=%v&p=%v", 1, 1), nil
}
	
func (j tx_catcher_quotes_data_job) ProcessResponse(ctx *base.FinDataCatcherJobContext, resp string) bool {
	//ctx
	//info *DailyDownloadeInfo
	
	// download_info := ctx.DownloadInfo

	// YY-MM-DD 前缀

	// 
	return false
}