package main

import "use_crontab/practice"

func main() {
	go practice.CreateServer()
	practice.StartTimer()

	//jsonData := `{"pre_order_id":1, "user_id":123,"op_type":100065,"app_id":301}`
	//processor := new (practice.Processor)
	//processor.ProcessOrder(jsonData)
	//defer practice.MyRecover()
	//retArr := practice.ReadData(3)
	//fmt.Printf("retArr:%#v \n", retArr)

	/*defer func() {
		r := recover()
		fmt.Printf("r:%#v \n", r)
		if  err1, ok := r.(error); ok {
			// 捕获异常
			fmt.Println("捕获了这个错误！", err1)
		} else {
			// 重新抛出一个异常
			panic(r)
		}
	}()

	err := errors.New(fmt.Sprintf("未查到该数据, pre_order_id:%d \n", 11))
	panic(err)*/
}
