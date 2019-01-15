package main

import(
	"fmt"
	"strings"

	"github.com/FinDataCatcher/base"
	"github.com/FinDataCatcher/utils"
	"github.com/FinDataCatcher/tx_catcher"

	"github.com/syndtr/goleveldb/leveldb"
)

var (
	DATA_DIR = "/home/ubuntu/FinDataCatcher.git/"
)

func TestTickDetail(str string) {
	ret, code, index, detail_list := tx_catcher.TX_parse_stock_tick_detail_response(str, "2019-01-10")

	fmt.Printf("===> %v\n", str)
	fmt.Printf("ret %v\n", ret)
	fmt.Printf("code %v\n", code.Detail())
	fmt.Printf("index %v\n", index)
	fmt.Printf("detail_list %v\n", detail_list)
	fmt.Printf("<===%v\n", str)
	
}

func TestQuotesDetail(str string) {
	ret, code, quotes := tx_catcher.TX_parse_stock_quotes_detail_response(str)

	fmt.Printf("===> %v\n", str)
	fmt.Printf("ret %v\n", ret)
	fmt.Printf("code %v\n", code.Detail())
	fmt.Printf("quotes %v\n", quotes)
	fmt.Printf("<===\n")
}



func TestLevelDB() {
	db, err := leveldb.OpenFile(fmt.Sprintf("%s/db_cache", DATA_DIR), nil)
	if nil != err {
		fmt.Printf("OpenDBCache file failed! %v\n", err)
		return
	}

	defer db.Close()

	data, err := db.Get([]byte("600102"), nil)
	if nil != err && err != leveldb.ErrNotFound { 
		fmt.Printf("Read file failed! Error : %v\n", err.Error())	
		return
	}

	fmt.Printf("Read config data %v! \n", data)

	info, err := base.DecodeDailyDownloadInfo(data)
	if nil != err {
		fmt.Printf("Decode daily download info failed %v! \n", err)		
		return
	}

	fmt.Printf("info %v\n", info)

	// Download Now
	//
	info.LastTickDataIndex++

	// Save
	data, err = base.EncodeDailyDownloadInfo(info)
	if nil != err {
		fmt.Printf("Encode daily download info failed %v! \n", err)		
		return	
	}
	
	err = db.Put([]byte("600102"), data, nil)
	if nil != err { 
		fmt.Printf("Write file failed! Error : %v\n", err.Error())	
		return
	}
}

func TestMysqlDB() {
	//  mysql 的Connector, 用来保存每组tick数据
	// 
}

func main() {
	TestTickDetail("v_detail_data_sz002451=[1,\"70/09:33:27/12.78/-0.01/251/321027/S|71/09:33:31/12.80/0.02/392/501495/B\"]")
	TestTickDetail("v_detail_data_sz300351=[0,\"0/09:25:04/9.87/-0.42/5415/5344605/S|1/09:30:04/9.85/-0.02/1210/1193398/S|2/09:30:07/9.84/-0.01/3258/3205853/S\"]")
	TestTickDetail("v_detail_data_sh600550=[0,\"0/09:25:02/4.07/0.01/282/114774/B|1/09:30:01/4.07/0.00/95/38690/S|2/09:30:04/4.07/0.00/573/233044/S\"]")

	TestLevelDB()


	fmt.Printf("%v\n", utils.CurrentLocalDate())
	fmt.Printf("%v\n", utils.CurrentLocalTime())

	// quoetes data
	v := strings.Split("1~zvzd~600519~358.74~361.29~361.88~27705~12252~15453~358.75~8~358.74~4~358.72~7~358.71~6~358.70~5~358.77~3~358.78~2~358.79~16~358.80~4~358.86~1~14:59:59/358.75/5/S/179381/28600|14:59:56/358.75/1/S/35875/28594|14:59:53/358.75/1/S/35875/28588|14:59:50/358.75/1/S/35875/28579|14:59:47/358.75/4/B/143499/28574|14:59:41/358.72/4/S/143501/28562~20170221150553~-2.55~-0.71~362.43~357.18~358.75/27705/994112865~27705~99411~0.22~27.24~~362.43~357.18~1.45~4506.49~4506.49~6.57~397.42~325.16~0.86", "~")
	fmt.Printf("Split Result %v\n", len(v))
	for index, s := range v {
		fmt.Printf("[%v] %v\n", index, s)	
	}

	// Test
	fmt.Printf("Show Timestamp %v %v\n", "20190115113456", tx_catcher.TX_ShowQuotesTimestamp("20190115113456"))

	// 
	TestQuotesDetail("v_sh600519=\"1~zvzd~600519~358.74~361.29~361.88~27705~12252~15453~358.75~8~358.74~4~358.72~7~358.71~6~358.70~5~358.77~3~358.78~2~358.79~16~358.80~4~358.86~1~14:59:59/358.75/5/S/179381/28600|14:59:56/358.75/1/S/35875/28594|14:59:53/358.75/1/S/35875/28588|14:59:50/358.75/1/S/35875/28579|14:59:47/358.75/4/B/143499/28574|14:59:41/358.72/4/S/143501/28562~20170221150553~-2.55~-0.71~362.43~357.18~358.75/27705/994112865~27705~99411~0.22~27.24~~362.43~357.18~1.45~4506.49~4506.49~6.57~397.42~325.16~0.86\";")

}
