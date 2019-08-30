package practice

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("inside hello()")
}

func anotherHello(T chan bool) {
	fmt.Println("inside anotherHello()")
	T <- true   //在协程方法的最后将true写入信道，以便等待主线程的接收，主线程如果不接收信道中的值，则会永远堵塞。
}

func NotUseChannel() {
	go hello()
	time.Sleep(50*time.Millisecond) //如果没有该语句，会看不到hello()的输出内容，因为主线程不会等待协程的运行结果
	fmt.Println("in NotUseChannel()")
}

func UseChannel() {
	channel := make(chan bool)
	go anotherHello(channel)
	data := <- channel	    //该语句是必需的，用于接收信道中的值，即使不赋值也是可以的。缺少的话，程序就会堵塞在这里,之后的语句不再执行。
	fmt.Printf("in UseChannel() data=%t\n", data)

	//fmt.Printf("in UseChannel() \n")
}
