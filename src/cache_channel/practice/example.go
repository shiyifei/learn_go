package practice

import (
	"fmt"
	"time"
)

/**
	在同一个方法中可以对缓存信道读和写
 */
func ReadWrite() {
	channel := make(chan string, 2)
	channel <- "专心致志"
	channel <- "爱拼才会赢"

	fmt.Println("receive from channel:", <-channel)
	fmt.Println("receive data from channel:", <-channel)
}

func write(channel chan int) {
	for i:=0;i<5;i++ {
		channel <- i
		fmt.Println("write to channel:",i)
	}
	close(channel)
}

/**
	测试缓存信道如何写入和读取
 */
func HowToCache() {
	channel := make(chan int, 2)  	//定义一个容量为2的缓存信道
	go write(channel)				//只要信道中的数据少于2个，该协程就会不断地写入数据，写满两个然后阻塞该协程
	time.Sleep(100*time.Millisecond)

	for num := range channel {
		fmt.Println("receive data from channel:", num)  //只要读取一个数就会影响导致信道中的实际元素的个数
		time.Sleep(100*time.Millisecond)
	}

}

