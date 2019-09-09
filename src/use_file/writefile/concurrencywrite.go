package writefile

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
)

/**
	生产者协程
 */
func produce(data chan int, wg *sync.WaitGroup) {
	n := rand.Intn(1000)
	data <- n
	wg.Done()
	fmt.Println("in produce()",n)
}

/**
	消费者协程
 */
func consume(data chan int, done chan bool) {
	f, err := os.Create("concurrent.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	for v := range data {
		_,err = fmt.Fprintln(f, v)
		if err != nil {
			fmt.Println(err)
			f.Close()
			done <- false
			return
		}
		fmt.Println("write file an number:",v)
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		done <- false
		return
	}
	done <- true
}

/**
	测试将100个随机数写入到同一个文件的示例，写入有并发控制，当前从通道中读取一个数写入。
 */
func Test() {
	data := make(chan int)    //初始化一个int类型的channel，里面存放100个随机数
	done := make(chan bool)		//初始化一个bool类型的channel,作为消费协程处理完毕的标志
	var wg sync.WaitGroup
	for i:=0;i<100;i++ {
		wg.Add(1)
		go produce(data, &wg)
	}

	go consume(data, done) //消费生产出的随机数，并写入完成状态到通道中

	go func(){
		fmt.Println("in func(),wg.Wait()")
		wg.Wait()     		//等待所有消费者协程处理完毕
		close(data)			//随机数channel用完之后关闭
	}()

	ret := <-done			//等待接收信道中的结果，程序将在这里阻塞
	if ret == true {
		fmt.Println("concurrent.txt have written successfully")
	} else {
		fmt.Println("concurrent.txt have written failed")
	}

}