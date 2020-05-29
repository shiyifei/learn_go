package main

import (
	"fmt"
	"os"
	"time"
	"use_rabbitmq/practice"
)

func main() {
	for i, v := range os.Args {
		fmt.Printf("args[%d]=%v \n", i ,v)
	}
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
