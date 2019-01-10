package main

import(
	"fmt"
	"github.com/FinDataCatcher/tx_catcher"
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

func main() {
	TestTickDetail("v_detail_data_sz002451=[1,\"70/09:33:27/12.78/-0.01/251/321027/S|71/09:33:31/12.80/0.02/392/501495/B\"]")
	TestTickDetail("v_detail_data_sz300351=[0,\"0/09:25:04/9.87/-0.42/5415/5344605/S|1/09:30:04/9.85/-0.02/1210/1193398/S|2/09:30:07/9.84/-0.01/3258/3205853/S\"]")
	TestTickDetail("v_detail_data_sh600550=[0,\"0/09:25:02/4.07/0.01/282/114774/B|1/09:30:01/4.07/0.00/95/38690/S|2/09:30:04/4.07/0.00/573/233044/S\"]")
}
