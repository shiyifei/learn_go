package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
)
var wg sync.WaitGroup


func main() {
	num := runtime.NumCPU()
	fmt.Println("Number of CPU is:",num)
	runtime.GOMAXPROCS(num)

	wg.Add(2)

	fmt.Println("Create Goroutines")
	go PrintPrime("A")
	go PrintPrime("B")

	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("Terminating Program")
}

//显示5000以内的素数
func PrintPrime(prefix string) {
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