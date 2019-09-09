package main

import (
	"time"
	channel "use_channel/practice"          //默认类名是practice,现在取别名为 channel1
)

func main() {
	channel.NotUseChannel()
	time.Sleep(100*time.Millisecond)  //100ms
	channel.UseChannel()


	//计算数字上的各位的平方和以及立方和，最后再求总和
	channel.PrintResult(987)

	//另一种写法：计算数字上的各位的平方和以及立方和，
	channel.ShowResult(987)

	//channel导致死锁，接收数据的写法，唯送信道相关
	channel.Test()

	//下面的程序会报错，在同一个协程中写入和读取信道，会导致阻塞死锁。
	// 向信道发送数据要放到其他goroutine中，读取可以放在本协程或其他协程中。
	/*chanTest := make(chan int)
	chanTest <- 666
	data := <- chanTest
	fmt.Println("receive from chanTest:", data)
	time.Sleep(100*time.Millisecond)*/
}

