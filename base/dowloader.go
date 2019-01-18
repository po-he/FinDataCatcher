package base

import(
	"errors"
	"encoding/gob"
	"bytes"
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
