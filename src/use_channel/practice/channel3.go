package practice

import (
	"fmt"
	"time"
)

/**
	信道是用于在不同的routine之间发送接收数据使用的，在一个方法中同时有发送和接收操作会不起作用。
 */
func receiveFromChan(chInt chan int) {
	fmt.Println("in receiveFromChan(),111")
	//chInt <- 888
	fmt.Println("in receiveFromChan(),222")   	//如果13行代码不注释掉的话该语句不会执行。
	fmt.Println("receive from chInt:", <-chInt)
	fmt.Println("in receiveFromChan(),333")
}

/**
	声明仅支持发送操作的信道： chSend chan<- int
 */
func sendData(chSend chan<- int) {
	fmt.Println("in SendOnlyChannel()")
	chSend <- 222
}

/**
	向channel中写入数据
 */
func sendToChannel(chInput chan int) {
	for i:=0; i<10;i++ {
		chInput <- i+1
	}
	close(chInput) //关闭信道
}

/**
	使用for循环来读取信道中的值
 */
func receiveFromChannel1(chInput chan int) {
	for {
		num, ok := <-chInput
		if ok == false {
			break
		}
		fmt.Printf("number is %d, is exists:%t \n", num, ok)
	}

}

/**
	使用range读取信道中的所有值
 */
func receiveFromChannel2(chInput chan int) {
	for num := range chInput {
		fmt.Println("number is ",num)
	}
}

func Test() {
	chInt := make(chan int)
	//chInt <- 666   //因为没有并发协程来读取这些信道，这里会出现死锁,出现错误：fatal error: all goroutines are asleep - deadlock!
	//time.Sleep(200*time.Millisecond)

	go sendData(chInt)   //一个协程发送数据
	//num := <-chInt       //主协程中接收数据的写法能够生效
	//fmt.Println("receive from chInt:", num)
	//time.Sleep(200*time.Millisecond)  //即使加入延时，另一个协程也接收不到数据
	//return
	go receiveFromChan(chInt) //一个协程接收，如果该方法对信道既发送又接收，会导致接收语句不生效。
	time.Sleep(200*time.Millisecond)


	//声明一个仅支持发送数据的信道
	chSend := make(chan<- int)
	go sendData(chSend)
	//number := <- chSend   //该语句会报错，因为不支持从仅支持发送数据的信道中读取值

	//声明一个支持发送以及接受数据的信道
	chReadWrite := make(chan int)
	go sendData(chReadWrite)   //形参如果是只写信道，传入读写信道是允许的。
	number := <- chReadWrite
	fmt.Println("read from read write channel, number=", number)
	time.Sleep(200*time.Millisecond)

	//测试两种循环读取信道中的值的方法
	go sendToChannel(chReadWrite)
	//go receiveFromChannel1(chReadWrite)
	go receiveFromChannel2(chReadWrite)
	time.Sleep(200*time.Millisecond)

}





