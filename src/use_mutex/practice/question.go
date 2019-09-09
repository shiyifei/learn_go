package practice

import (
	"fmt"
	"sync"
)

var x int = 1

/**
	没加锁情况下会出现竞态条件的情况，导致最终结果与预期不一致。
 */
func increment(wg * sync.WaitGroup) {
	x = x+1
	wg.Done()
}

/**
	加入排他锁之后，任何时候只允许一个协程进入到该函数内部执行赋值操作。
 */
func incrementByMutex(wg * sync.WaitGroup, mutex * sync.Mutex) {
	mutex.Lock()
	x = x+1
	mutex.Unlock()
	wg.Done()
}

/**
	显示排它锁在使用和不使用的情况下程序的最终输出结果
 */
func Display() {
	var wg sync.WaitGroup
	for i:=0;i<100;i++ {
		wg.Add(1)
		go increment(&wg)
	}
	wg.Wait()
	fmt.Println("not use mutex, final value of x is :",x)	//最终结果是不确定的

	var mutex sync.Mutex
	x=1
	for i:=0;i<100;i++ {
		wg.Add(1)
		go incrementByMutex(&wg, &mutex)
	}
	wg.Wait()
	fmt.Println("use mutex,final value of x is :",x)  	//最终结果一定是101

}