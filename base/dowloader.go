package base

import(
	"errors"
	"encoding/gob"
	"bytes"
	"fmt"
	"time"
	"net/http"
	"io/ioutil"
)

//// TODO : 
// 1. multi-goroutine 同時寫日志會不會交錯
// 

var (
	DEFAULT_TIMEOUT_SECOND = 15
)

type DailyDownloadeInfo struct {
	//Date  
	CodeStr  string
	LastTickDataIndex int32

}

func EncodeDailyDownloadInfo(info *DailyDownloadeInfo) ([]byte, error) {
	if nil == info {
		return []byte{}, errors.New("Nil DailyDownloadeInfo")
	}

	buffer := new(bytes.Buffer)
	enc := gob.NewEncoder(buffer)
	if err := enc.Encode(info); err != nil {
		return []byte{}, err
	}

	return buffer.Bytes(), nil
}

func DecodeDailyDownloadInfo(data []byte) (*DailyDownloadeInfo, error) {
	info := &DailyDownloadeInfo{}
	if len(data) == 0 {
		return info, nil
	}
	
	enc := gob.NewDecoder(bytes.NewReader(data))
	if err := enc.Decode(info); err != nil {
		return nil, err
	} else {
		return info, nil	
	}
}

func DoHttpRequest(url string, timeout int) ([]byte, error) {
	if 0 == timeout {
		timeout = DEFAULT_TIMEOUT_SECOND
	}

	c := &http.Client{
				Timeout : time.Duration(timeout) * time.Second,
		  }

	httpResp, err := c.Get(url)
	defer func() {
		if nil != httpResp {
			httpResp.Body.Close()	
		}
	} ()

	if nil != err {
		return []byte{}, err		
	}

	body, err := ioutil.ReadAll(httpResp.Body)
	if nil != err {
		return []byte{}, err		
	}

	return body, nil
}

type DataCatcherJobOperator interface {
	RequestData(*FinDataCatcherJobContext)  (string, error)	
	ProcessResponse(*FinDataCatcherJobContext, string) error
}

type FinDataCatcherJobContext struct {
	DownloadInfo *DailyDownloadeInfo
	Job  DataCatcherJobOperator
	Url  string

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
