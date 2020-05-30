/**
	一种修正代码、消除竞争状态的办法是,使用Go语言提供的锁机制，锁住共享资源从而保证goroutine的同步状态
	对于共享资源的读写操作是必须是原子性的，同一个时刻只有一个goroutine对共享资源进行读写操作才合适
	Go语言中提供了原子访问atomic包, "sync/atomic" ,原子函数能够以很底层的加锁机制来同步访问整型变量和指针
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

//注意本示例，如果取消注释行会导致有数据竟态
//使用 go run -race use_goroutine.go 可以看到结果
func increCounterNew(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	for count:=0; count<2; count++ {
		//fmt.Println("current routine id:",id,",before assign,count=",count,",counterNew=",counterNew)
		fmt.Println("current routine id:",id,",before assign,count=",count,",counterNew=",atomic.LoadInt64(&counterNew))
		//安全地对counterNew加1,强制同一时刻只能有一个goroutine运行并完成这个加法操作
		atomic.AddInt64(&counterNew, 1)

		//fmt.Println("current routine id:",id,",after assign,count=",count,",counterNew=",counterNew)  该语句会导致两个协程同时读取一个全局变量，并发读取有可能导致读取结果错误。
		//替换成 atomic.LoadInt64(&counterNew) 就能避免竟态情况的产生
		fmt.Println("current routine id:",id,",after assign,count=",count,",counterNew=", atomic.LoadInt64(&counterNew))


		//当前goroutine从线程退出,并放回到队列
		runtime.Gosched()
		fmt.Println("current routine id:",id,",count=",count,",counterNew=", atomic.LoadInt64(&counterNew))
	}
}

