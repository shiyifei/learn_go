package main

import (
	"time"
	"use_rabbitmq/practice"
)

func main() {
	//chanRet := make(chan bool)
	//chanRet <-
	//
	practice.Produce()
	//
	//result := <-chanRet
	time.Sleep(300* time.Millisecond)
	practice.Consume()

	//practice.GenerateTask([]string{"aaa","bbb","ccc", "ddd"})
	//
	//practice.ProcessTask()
	//practice.ProcessTask()

}
