package practice

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	num int64
	wgTest sync.WaitGroup
)

func Do(name string) {
	defer wgTest.Done()

	for {
		fmt.Println("name=",name)
		time.Sleep(100*time.Millisecond)
		if 1 == atomic.LoadInt64(&num) {
			fmt.Println("goroutine is stop", name)
			break
		}
	}
}

func AtomicLoadStore() {
	wgTest.Add(2)
	go Do("zhangsan")
	go Do("lisi")
	time.Sleep(500*time.Millisecond)
	fmt.Println("I hope goroutine stop")
	atomic.StoreInt64(&num, 1)
	wgTest.Wait()
}