/**
	一种修正代码、消除竞争状态的办法是,使用Go语言提供的锁机制，锁住共享资源从而保证goroutine的同步状态
 */
package practice

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	counterNew int64
)

func Practice4() {
	var wg sync.WaitGroup
	wg.Add(2)

	go increCounterNew(1, &wg)
	go increCounterNew(2, &wg)

	wg.Wait()
	fmt.Println("Final Counter:",counterNew)
}

func increCounterNew(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	for count:=0; count<2; count++ {
		fmt.Println("current routine id:",id,",before assign,count=",count,",counterNew=",counterNew)
		//安全地对counterNew加1,强制同一时刻只能有一个goroutine运行并完成这个加法操作
		atomic.AddInt64(&counterNew, 1)

		fmt.Println("current routine id:",id,",after assign,count=",count,",counterNew=",counterNew)

		//当前goroutine从线程退出,并放回到队列
		runtime.Gosched()
		fmt.Println("current routine id:",id,",count=",count,",counterNew=",counterNew)
	}
}

