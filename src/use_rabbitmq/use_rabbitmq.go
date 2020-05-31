package main

import (
	"fmt"
	"os"
	"time"
	"use_rabbitmq/practice"
)

func main() {
	/*for i, v := range os.Args {
		fmt.Printf("args[%d]=%v \n", i ,v)
	}*/

	defer practice.SqlDB.Close()
	var action string
	if len(os.Args)>1 {
		action = string(os.Args[1])
	}
	if action == "product" {
		//practice.Produce()
		practice.MultiSendMsg()
		return
	} else if action == "consume" {
		practice.MultiConsume()
		return
	} else if action == "testpost" {
		ret, err := practice.SendPostRequest("http://192.168.1.102:8100/mq/consume", "{\"name\":\"user490\",\"addr\":\"Beijing\",\"age\":34}")
		fmt.Println("ret:",ret,"err:", err)
		return
	}

	//practice.Produce()
	practice.MultiSendMsg()
	//fmt.Println("=========== after sending message ==================================")
	time.Sleep(500* time.Millisecond)
	//practice.Consume()

	practice.MultiConsume()



	//practice.GenerateTask([]string{"aaa","bbb","ccc", "ddd"})
	//
	//practice.ProcessTask()
	//practice.ProcessTask()

}
