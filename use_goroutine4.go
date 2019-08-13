/**
	一种修正代码、消除竞争状态的办法是,使用Go语言提供的锁机制，锁住共享资源从而保证goroutine的同步状态
 */
package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	counter int64

	wg sync.WaitGroup
)

func main() {
	wg.Add(2)

	go increCounter(1)
	go increCounter(2)

	wg.Wait()
	fmt.Println("Final Counter:",counter)
}

func increCounter(id int) {
	defer wg.Done()

	for count:=0; count<2; count++ {
		fmt.Println("current routine id:",id,",before assign,count=",count,",counter=",counter)
		//安全地对counter加1,强制同一时刻只能有一个goroutine运行并完成这个加法操作
		atomic.AddInt64(&counter, 1)

		fmt.Println("current routine id:",id,",after assign,count=",count,",counter=",counter)

		//当前goroutine从线程退出,并放回到队列
		runtime.Gosched()
		fmt.Println("current routine id:",id,",count=",count,",counter=",counter)
	}
}

