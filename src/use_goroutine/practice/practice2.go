package practice

import (
	"fmt"
	"math"
	"runtime"
	"sync"
)


func Practice2() {
	num := runtime.NumCPU()
	fmt.Println("Number of CPU is:",num)
	runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Create Goroutines")
	go PrintPrime("A", &wg)
	go PrintPrime("B", &wg)

	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("Terminating Program")
}

/**
	显示5000以内的素数,查找并显示素数会消耗不少时间，
	这会让调度去有机会在第一个goroutine找到所有素数之前，切换该goroutine的时间片
 */
func PrintPrime(prefix string, wg * sync.WaitGroup) {
	defer wg.Done()

	next:
		for outer:=2; outer<5000; outer++ {
			for inner:=2; inner <= int(math.Floor(math.Sqrt(float64(outer)))); inner++ {
				if outer%inner == 0 {
					continue next
				}
			}
			fmt.Printf("%s:%d \t", prefix, outer)
		}
		fmt.Println("Completed", prefix)
}

//当goruntine方法执行时间较长时，会发现可能有相互穿插执行的情况，时间片轮询占用逻辑处理器