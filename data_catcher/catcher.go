package data_catcher

import(
	"fmt"

	"github.com/FinDataCatcher/base"
)


type DataCatcherJobOperator interface {
	RequestData(*FinDataCatcherJobContext)  (string, error)	
	ProcessResponse(*FinDataCatcherJobContext, string) bool
}

type FinDataCatcherJobContext struct {
	DownloadInfo *base.DailyDownloadeInfo
	Job  DataCatcherJobOperator
	Url  string

	MarketCloseTimestamp    int64     // Job 
	// TODO : 增加BerkleyDB
	// TODO : 增加MysqlDB 的 Connect
}

func (ctx *FinDataCatcherJobContext) SaveDailyDownloadInfo() {
	/*
	db, err := leveldb.OpenFile(fmt.Sprintf("%s/db_cache", DATA_DIR), nil)
	if nil != err {
		fmt.Printf("OpenDBCache file failed! %v\n", err)
		return
	}

	defer db.Close()

	info := ctx.DownloadInfo
	// Save
	data, err = base.EncodeDailyDownloadInfo(info)
	if nil != err {
		fmt.Printf("Encode daily download info failed %v! \n", err)		
		return	
	}
	
	err = db.Put([]byte(info), data, nil)
	if nil != err { 
		fmt.Printf("Write file failed! Error : %v\n", err.Error())	
		return
	}		
	*/
} 

type FinDataHttpCatcher struct {
	dispatchChan chan *FinDataCatcherJobContext
	Id           int

	bind_ctx *FinDataCatcherJobContext
}

func CreateFinDataHttpCatcher() *FinDataHttpCatcher {
	c := FinDataHttpCatcher {
			dispatchChan : make(chan *FinDataCatcherJobContext, 1),			
		 }

	return &c		 
}

func (catcher *FinDataHttpCatcher) doJob() {
	bind_ctx := catcher.bind_ctx
	resp, err := bind_ctx.Job.RequestData(bind_ctx)
	if nil != err {
		fmt.Printf("")
		return
	}
	
	bind_ctx.Job.ProcessResponse(bind_ctx, resp)	
}

func (catcher *FinDataHttpCatcher) AssignJobContext(ctx *FinDataCatcherJobContext) {
	if nil == ctx {
		return
	}

	catcher.dispatchChan <- ctx	
}

func (catcher *FinDataHttpCatcher) Idle() {
	// 把自己放回調度隊列末尾
	d := GetGlobalFinDataCatcherMng()
	if nil == d {
		return
	}

	d.AddHttpCatcher(catcher)
}

func (catcher *FinDataHttpCatcher) Run() {
	// defer 

	for {
		select {
			case ctx, ok := <-catcher.dispatchChan:
			if ok {
				//gs_users.UpdateSavingUserList(uList)
				catcher.bind_ctx = ctx
				catcher.doJob()
				catcher.Idle()
			}	
		}
	}	
}
