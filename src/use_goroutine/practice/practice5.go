/**
	使用互斥锁(mutex)解决两个goroutine引发数据竞争的情况，竞争状态的存在让并发程序变得复杂，十分容易引起潜在问题。
 */
package practice

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	// counter1 是所有goroutine都要增加其值的变量
	counter1 int

	//wg1 用来等待程序结束
	wg1 sync.WaitGroup

	//mutex 用来定义一段代码临界区
	mutex sync.Mutex
)

func Practice5() {
	wg1.Add(2)

	go increCounter(1)
	go increCounter(2)

	wg1.Wait()
	fmt.Println("Final Counter:", counter1)
}

func increCounter(id int) {
	defer wg1.Done()
	fmt.Println("current routine id:", id)


	for count:=0; count<2; count++ {
		//同一时刻只允许一个goroutine进入这个临界区
		mutex.Lock()
		{
			value := counter1
			fmt.Println("current routine id:",id,",count=",count,",value=",value)

			//当前goroutine从线程退出,并放回到队列,让其他goroutine运行
			//强制将当前goroutine退出当前线程后，调度器会再次分配这个goroutine继续运行。即：调度器强制不切换到另一个routine
			runtime.Gosched()
			fmt.Println("current routine id:", id, ",after Gosched(),counter1=", counter1)
			value++
			counter1 = value

			fmt.Println("current routine id:", id, ",after assign,counter1=", counter1)
		}

		//释放锁，允许其他正在等待的goroutine进入临界区
		mutex.Unlock()
	}
}