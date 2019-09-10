package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	begin := time.Now()
	//分配一个逻辑处理器给调度器使用,
	runtime.GOMAXPROCS(1)
	//runtime.GOMAXPROCS(runtime.NumCPU())

	//wg用来等待程序完成 waitGroup是一个计数信号量，用来记录并维护运行的goroutine
	//计算加2,表示要等待两个goroutine
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start Goroutines")

	// 声明一个匿名函数，并创建一个goruntine
	go func() {
		//在函数退出时，调用Done通知main函数工作已经完成，Done()方法会减小WaitGroup的值
		defer wg.Done()

		//显示字母表三次
		for count:=0; count<300; count++ {
			for char := 'a'; char <'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	go func() {
		defer wg.Done()

		for count:=0;count<300; count++ {
			for char:='A';char<'A'+26; char++ {
				fmt.Printf("%c ",char)
			}
		}
	}()

	//等待goroutine结束
	fmt.Println("Waiting To Finish")
	//如果WaitGroup的值大于0，则Wait()方法就会阻塞。
	wg.Wait()

	end := time.Now()
	diff := end.Sub(begin)
	fmt.Println("process this data, time interval:", diff)

	fmt.Println("\n Terminating Program")
}
