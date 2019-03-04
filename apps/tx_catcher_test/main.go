package main

import(
	"fmt"
	"sync"
	"time"
	"math/rand"
	
	//"string"

	"github.com/FinDataCatcher/base"
	"github.com/FinDataCatcher/tx_catcher"
	"github.com/FinDataCatcher/cache"
	"github.com/FinDataCatcher/db/mongodb"
	
	//"github.com/syndtr/goleveldb/leveldb"
	"github.com/FinDataCatcher/utils"
	
)

var (
	//DATA_DIR = "/home/ubuntu/FinDataCatcher.git/"
	DATA_DIR = "/home/heyang/FinDataCatcher.git/"
	DATA_DIR_2 = "/home/heyang/testCache/"


	tickInfoDownloadCache *cache.Cache

	_wg sync.WaitGroup 
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



// func TestLevelDB() {
// 	db, err := leveldb.OpenFile(fmt.Sprintf("%s/db_cache", DATA_DIR), nil)
// 	if nil != err {
// 		fmt.Printf("OpenDBCache file failed! %v\n", err)
// 		return
// 	}

// 	defer db.Close()

// 	data, err := db.Get([]byte("600102"), nil)
// 	if nil != err && err != leveldb.ErrNotFound { 
// 		fmt.Printf("Read file failed! Error : %v\n", err.Error())	
// 		return
// 	}

// 	fmt.Printf("Read config data %v! \n", data)

// 	info, err := base.DecodeDailyDownloadInfo(data)
// 	if nil != err {
// 		fmt.Printf("Decode daily download info failed %v! \n", err)		
// 		return
// 	}

// 	fmt.Printf("info %v\n", info)

// 	// Download Now
// 	//
// 	info.LastTickDataIndex++

// 	// Save
// 	data, err = base.EncodeDailyDownloadInfo(info)
// 	if nil != err {
// 		fmt.Printf("Encode daily download info failed %v! \n", err)		
// 		return	
// 	}
	
// 	err = db.Put([]byte("600102"), data, nil)
// 	if nil != err { 
// 		fmt.Printf("Write file failed! Error : %v\n", err.Error())	
// 		return
// 	}
// }

func TestMysqlDB() {
	//  mysql 的Connector, 用来保存每组tick数据
	// 
}

func TestCache(str string) {

	data, err := tickInfoDownloadCache.Read(str)
	if nil != err {
		fmt.Printf("Read info failed %v! \n", err)		
		return
	}


	info, err := base.DecodeDailyDownloadInfo(data)
	if nil != err {
		fmt.Printf("Decode daily download info failed %v! \n", err)		
		return
	}

	fmt.Printf("info %v\n", info)

	// Download Now
	if "" == info.CodeStr {
		info.CodeStr = str
	}

	info.LastTickDataIndex++

// 	// Save
	data, err = base.EncodeDailyDownloadInfo(info)
	if nil != err {
		fmt.Printf("Encode daily download info failed %v! \n", err)		
		return	
	}

	tickInfoDownloadCache.Write(str, data)

	_wg.Done()
}

func TestInitTickInfoDownloadCache() {
	var err error
	tickInfoDownloadCache, err = cache.CreateCacheObj(DATA_DIR_2)
	if nil != err {
		fmt.Printf("CreateNewCacheObj failed! path[%v] %v\n", DATA_DIR_2, err)
		return
	} else if err = tickInfoDownloadCache.Open(); nil != err {
		fmt.Printf("Open TickInfoDownloadCache failed! %v\n",  err)
		return
	}
}

func TestCloseTickInfoDownloadCache() {
	tickInfoDownloadCache.Close()
}



func main() {
	//TestTickDetail("v_detail_data_sz002451=[1,\"70/09:33:27/12.78/-0.01/251/321027/S|71/09:33:31/12.80/0.02/392/501495/B\"]")
	//TestTickDetail("v_detail_data_sz300351=[0,\"0/09:25:04/9.87/-0.42/5415/5344605/S|1/09:30:04/9.85/-0.02/1210/1193398/S|2/09:30:07/9.84/-0.01/3258/3205853/S\"]")
	//TestTickDetail("v_detail_data_sh600550=[0,\"0/09:25:02/4.07/0.01/282/114774/B|1/09:30:01/4.07/0.00/95/38690/S|2/09:30:04/4.07/0.00/573/233044/S\"]")

	// TestLevelDB()


	//fmt.Printf("%v\n", utils.CurrentLocalDate())
	//fmt.Printf("%v\n", utils.CurrentLocalTime())

	// quoetes data
	// v := strings.Split("1~zvzd~600519~358.74~361.29~361.88~27705~12252~15453~358.75~8~358.74~4~358.72~7~358.71~6~358.70~5~358.77~3~358.78~2~358.79~16~358.80~4~358.86~1~14:59:59/358.75/5/S/179381/28600|14:59:56/358.75/1/S/35875/28594|14:59:53/358.75/1/S/35875/28588|14:59:50/358.75/1/S/35875/28579|14:59:47/358.75/4/B/143499/28574|14:59:41/358.72/4/S/143501/28562~20170221150553~-2.55~-0.71~362.43~357.18~358.75/27705/994112865~27705~99411~0.22~27.24~~362.43~357.18~1.45~4506.49~4506.49~6.57~397.42~325.16~0.86", "~")
	// fmt.Printf("Split Result %v\n", len(v))
	// for index, s := range v {
	// 	fmt.Printf("[%v] %v\n", index, s)	
	// }

	// // Test
	// fmt.Printf("Show Timestamp %v %v\n", "20190115113456", tx_catcher.TX_ShowQuotesTimestamp("20190115113456"))

	// // 
	// TestQuotesDetail("v_sh600519=\"1~zvzd~600519~358.74~361.29~361.88~27705~12252~15453~358.75~8~358.74~4~358.72~7~358.71~6~358.70~5~358.77~3~358.78~2~358.79~16~358.80~4~358.86~1~14:59:59/358.75/5/S/179381/28600|14:59:56/358.75/1/S/35875/28594|14:59:53/358.75/1/S/35875/28588|14:59:50/358.75/1/S/35875/28579|14:59:47/358.75/4/B/143499/28574|14:59:41/358.72/4/S/143501/28562~20170221150553~-2.55~-0.71~362.43~357.18~358.75/27705/994112865~27705~99411~0.22~27.24~~362.43~357.18~1.45~4506.49~4506.49~6.57~397.42~325.16~0.86\";")
	

	// TestInitTickInfoDownloadCache()

	// for i:=0; i<20; i++ {
	// 	_wg.Add(1)	
	// 	go TestCache(fmt.Sprintf("%d", 600102 + i))
	// }
	
	// _wg.Wait()
	
	// TestCloseTickInfoDownloadCache()


	// // 
	// TestUpsertMongoTimeSeriesData()

	//TestGenRandomNum()
	CheckHotZone()
}


/*
测试 mongodb 存储 时间序列数据
参考：https://www.mongodb.com/blog/post/schema-design-for-time-series-data-in-mongodb
*/
func TestInsertMongoTimeSeriesData() {
	// 
	/*> db.metrics.insert({"timestamp":ISODate("2019-02-16T17:00:00Z"), "type":"tick", "values":{0:{1:1, 2:2}}})
	 WriteResult({ "nInserted" : 1 })
	 > db.metrics.find()
	 { "_id" : ObjectId("5c67df88ad82e2bb37fccbde"), "timestamp" : ISODate("2019-02-16T17:00:00Z"), "type" : "tick", "values" : { "0" : { "1" : 1, "2" : 2 } } }
	 > db.metrics.update({"timestamp":ISODate("2019-02-16T17:00:00Z"), "type":"tick", "values.0.0":0})
	 > db.metrics.update({"timestamp":ISODate("2019-02-16T17:00:00Z"), "type":"tick"}, {$set: {"values.1.0":10}})
	   WriteResult({ "nMatched" : 1, "nUpserted" : 0, "nModified" : 1 })
	 > db.metrics.update({"timestamp":ISODate("2019-02-16T17:00:00Z"), "type":"tick"}, {$set: {"values.1.1":11}})
	   WriteResult({ "nMatched" : 1, "nUpserted" : 0, "nModified" : 1 })
	 > db.metrics.find()
	   { "_id" : ObjectId("5c67df88ad82e2bb37fccbde"), "timestamp" : ISODate("2019-02-16T17:00:00Z"), "type" : "tick", "values" : { "0" : { "1" : 1, "2" : 2, "0" : 0 }, "1" : { "0" : 10, "1" : 11 } } }
	*/
}

func TestUpsertMongoTimeSeriesData() {
	quotesDb := new(mongodb.QuotesMgoDatabase)
	err := quotesDb.Open("192.168.1.50", "", "", "test_quotes") 
	if nil != err {
		fmt.Printf("Open test_quotes failed!")
		return
	}

	defer quotesDb.Close()
	begin_sec := utils.NormalDate2LocalUnixTimeSec("2016-01-01 00:00:00")
	end_sec := utils.NormalDate2LocalUnixTimeSec("2016-01-31 00:00:00")
	for sec := begin_sec; sec <= end_sec; sec += int64(utils.ONE_DAY_SEC) {
		TestInsertQuotesData(quotesDb, sec) 
	}
}

func TestInsertQuotesData(db *mongodb.QuotesMgoDatabase, dayMidNightSec int64) {
	m_begin := dayMidNightSec + 9 * int64(utils.ONE_HOUR_SEC) + 30 * int64(utils.ONE_MINUTE)
	m_end   := dayMidNightSec + 11 * int64(utils.ONE_HOUR_SEC) + 30 * int64(utils.ONE_MINUTE)
	a_begin := dayMidNightSec + 13 * int64(utils.ONE_HOUR_SEC)
	a_end := dayMidNightSec + 15 * int64(utils.ONE_HOUR_SEC) 

	//
	cnt := int32(0)
    var err error
	for sec := m_begin; sec <= m_end; sec+=3 {
		// 
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
    	amount := int32(r.Intn(10000))
    	price := r.Float64()
		err = db.Insert(int64(sec), price, amount, 2) 
		if nil != err {
			panic("Insert morning QuotesFailed\n")
		}

		err = db.InsertSingleLine(int64(sec), price, amount, 2)
		if nil != err {
			panic("Insert morning SingleQuotesFailed\n")
		}

		cnt++
		if cnt >= 25 {
			time.Sleep(time.Duration(2)*time.Second)	
			cnt = 0
		}		
	}

	cnt = 0
	for sec := a_begin; sec <= a_end; sec+=3 { 
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
    	amount := int32(r.Intn(10000))
    	price := r.Float64()
		err = db.Insert(int64(sec), price, amount, 2) 
		if nil != err {
			panic("Insert afternoon QuotesFailed\n")
		}

		err = db.InsertSingleLine(int64(sec), price, amount, 2)
		if nil != err {
			panic("Insert afternoon SingleQuotesFailed\n")
		}

		cnt++
		if cnt >= 25 {
			time.Sleep(time.Duration(1)*time.Second)	
			cnt = 0
		}
	}

	fmt.Printf("-> %v finish! \n", utils.Sec2LocalDataString(dayMidNightSec))
}

func TestMongoTimeSeariesData() {
	//
}

func TestGenRandomNum() {
	l := make([]float64, 0)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
    val := float64(10.0)
	for i:=0; i<30; i++ {
		dis := int32(r.Intn(10))
		side := int32(r.Intn(2))
		if side == 0 {
			val = val * (float64(1.0) + (float64(dis) / 100))
		} else {
			val = val * (float64(1.0) - (float64(dis) / 100))
		}

		l = append(l, val)
	}

	str := ""
	for index, v := range l {
		if 0 == index {
			str = fmt.Sprintf("%.2f", v)
		} else {
			str = fmt.Sprintf("%v, %.2f", str, v)
		} 
	}

	fmt.Printf("%v \n", str)
}

var (
	debugMap map[string]bool
)


func CheckHotZone() {
	//v_list := []float64{10.00, 10.00, 9.90, 10.69, 10.59, 10.27, 10.16, 10.98, 10.65, 10.76, 9.89, 10.59, 10.80, 9.94, 10.33, 9.51, 9.60, 9.41, 9.88, 9.88, 9.48, 9.77, 9.38, 9.75, 10.05, 9.64, 10.13, 9.92, 10.42, 11.25}
	v_list := make([]float64, 0)
	for i:=0; i<30; i++ {
		v_list = append(v_list, 10.00)	
	}
	v_list_max_index := len(v_list) - 1

	debugMap = make(map[string]bool)


	for index, v := range v_list {
		fmt.Printf("[%v]%v ", index, v)
	}

	fmt.Printf("\n")
	forward_stage := int(5)
	index:=0
	for {
		if index >= v_list_max_index {
			break;
		}

		v :=  v_list[index] 
		needNextBegin := false
		max := v * float64(1+0.1)
		beginIndex := index

		for {
			endPos, find := searchHotEndIndex(v_list, beginIndex, forward_stage, v_list_max_index, max)
			if find == true {
				if endPos >= beginIndex {
					beginIndex = endPos
					//fmt.Printf("change max %v => %v \n", max, v_list[beginIndex])
					max = v_list[beginIndex]
					continue
				} 
			} else if find == false {
				if beginIndex > index {
					fmt.Printf("[%v/%v => %v/%v]\n", index, v_list[index], beginIndex, v_list[beginIndex])
					index = beginIndex
					needNextBegin = true
				} else {
					index += 1		
				}

				break;
			}	
		}

		if needNextBegin == false {
			continue
		}

		// 找到下一个最小值的开始位置
		min := v_list[index] * float64(1-0.1)
		beginIndex = index
		for {
			endPos, find := searchNextHotBeginIndex(v_list, beginIndex, forward_stage, v_list_max_index, min)
			if find == true {
				if endPos >= beginIndex {
					beginIndex = endPos
					min = v_list[beginIndex]
					continue
				} 
			} else if find == false {
				if beginIndex > index {
					index = beginIndex
					//fmt.Printf("[%v/%v => %v/%v]\n", )
					//needNextBegin = true
				} else {
					index += 1
				}

				//fmt.Printf("next begin pos %v\n", index)
				break;
			}	
		}
	}	
}

func searchHotEndIndex(v_list []float64, beginIndex, forward_stage, v_list_max_index int, max float64) (int, bool) {
	endIndex := beginIndex + forward_stage
	if endIndex >= v_list_max_index {
		endIndex = v_list_max_index
	} 

	//fmt.Printf("searchHotEndIndex %v => %v\n", beginIndex, endIndex)
	
	if len(v_list) == 0 || endIndex < beginIndex{
		return beginIndex, false
	}

	key := fmt.Sprintf("%v|%v", beginIndex, endIndex)
	if _, exist := debugMap[key]; exist {
		panic("panic")
	} else {
		debugMap[key] = true
	}


	maxVal := max
	maxValPos := beginIndex
	bFind := false
	for i:=beginIndex; i < v_list_max_index; i++ {
		if v_list[i] > maxVal {
			maxValPos = i
			maxVal = v_list[i] 
			bFind = true
		}
	}

	if bFind {
		return maxValPos, true
	} else {
		return beginIndex, false
	}
}

func searchNextHotBeginIndex(v_list []float64, beginIndex, forward_stage, v_list_max_index int, min float64) (int, bool) {
	endIndex := beginIndex + forward_stage
	if endIndex >= v_list_max_index {
		endIndex = v_list_max_index
	} 

	//fmt.Printf("searchHotBeginIndex %v => %v\n", beginIndex, endIndex)
	if len(v_list) == 0 || endIndex < beginIndex{
		return beginIndex, false
	}

	
	minVal := min
	minValPos := beginIndex
	bFind := false
	for i:=beginIndex; i < v_list_max_index; i++ {
		if v_list[i] < minVal {
			minValPos = i
			minVal = v_list[i] 
			bFind = true
		}
	}

	if bFind {
		return minValPos, true
	} else {
		return beginIndex, false
	}
}


func CheckEndSentinelPos(l []float64, max float64, maxoffset int) {

}






