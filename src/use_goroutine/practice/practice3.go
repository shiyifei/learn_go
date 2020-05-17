/**
	演示两个goroutine引发数据竞争的情况，竞争状态的存在让并发程序变得复杂，十分容易引起潜在问题。
 */
package practice

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	// counter 是所有goroutine都要增加其值的变量
	counter int
)

func Practice3() {
	var wg sync.WaitGroup  //wg 用来等待程序结束
	wg.Add(2)

	go incCounter(1, &wg)
	go incCounter(2, &wg)

	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

func incCounter(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("current routine id:", id)


	for count:=0; count<2; count++ {
		value := counter

		fmt.Println("current routine id:",id,",count=",count,",value=",value)
		//Gosched让出处理器，当前goroutine从线程退出,并放回到队列,让其他goroutine运行
		runtime.Gosched()

		value++
		counter = value

		fmt.Println("current routine id:",id,",counter=",counter)
	}
}