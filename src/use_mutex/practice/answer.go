package practice

import (
	"fmt"
	"sync"
)

var input int =1

/**
	使用缓冲信道来处理竞态条件的问题, 缓冲信道用于保证只有一个协程访问增加 input 的临界区
 */
func incrementByChannel(wg *sync.WaitGroup, ch chan int) {
	ch <- 1  	//信道中写入一个值后会阻塞其他协程不允许再写入该信道
	input = input+1
	<-ch		//赋值完毕后读取信道的值，这时候信道的长度又变为0，还可以继续写入值了。
	wg.Done()
}


func Test() {
	var wg sync.WaitGroup
	channel := make(chan int, 1)  //定义一个容量为1的信道
	for i:=0;i<100;i++ {
		wg.Add(1)
		go incrementByChannel(&wg, channel)
	}
	wg.Wait()
	fmt.Println("final value of input is:",input) 	//最终的输出结果也是101
}