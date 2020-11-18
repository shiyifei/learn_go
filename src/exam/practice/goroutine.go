package practice

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func printChar(times int, begin, end byte) {
	for count:=0; count<times; count++ {
		for char := begin; char <= end; char++ {
			fmt.Printf("%c", char)
		}
	}
}

var times int
func init() {
	times = 1000
}

func concurrency() {
	begin := time.Now()
	runtime.GOMAXPROCS(2)

	var wg sync.WaitGroup

	wg.Add(3)

	//打印100次小写字母
	go func() {
		defer wg.Done()
		printChar(times, 'a', 'z')
	}()

	//打印100次大写字母
	go func() {
		defer wg.Done()
		printChar(times, 'A', 'Z')
	}()

	//打印100次数字
	go func() {
		defer wg.Done()
		printChar(times, '0', '9')
	}()
	wg.Wait()

	diff :=  time.Now().Sub(begin)
	fmt.Println("\n 并发打印完所有字母，耗时：",diff)
}

func serial() {
	begin := time.Now()
	printChar(times, 'a', 'z')
	printChar(times, 'A', 'Z')
	printChar(times, '0', '9')
	diff :=  time.Now().Sub(begin)
	fmt.Println("\n 串行打印完所有字母，耗时：",diff)
}

func TestConcurrency() {
	concurrency()
	serial()
}

