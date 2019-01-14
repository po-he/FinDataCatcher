package main

import(
	"fmt"

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
}
