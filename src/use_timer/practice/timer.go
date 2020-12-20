package practice

import (
	"fmt"
	"reflect"
	"sync"
	"time"
)
const(
	//golang时间格式 加入毫秒显示 如果是6个0，则是微秒，3个0表示不足三位时左侧会补齐0
	TEMPLATE = "2006-01-02 15:04:05.000"
)

var wg sync.WaitGroup


func TestTimer() {
	wg.Add(2)

	timer := time.NewTimer(2*time.Second)
	go func(t *time.Timer){
		defer wg.Done()
		<- t.C
		fmt.Println("get timer, current time is:",time.Now().Format(TEMPLATE))

		//Reset使t重新开始计时，本方法返回后再等待一段时间d后到期。
		//如果调用Reset方法时，t还在等待中会返回真；t已经到期或停止了，会返回假
		ret := t.Reset(2*time.Second)
		fmt.Println("type of ret:", reflect.TypeOf(ret),"return:",ret)
	}(timer)



	//NewTicker 返回一个新的Ticker,该Ticker包含一个通道字段，并会每隔时间段d后就向该通道发送当时的时间。
	//它会调整时间间隔或者丢弃tick信息以适应慢的接受者 如果d<=0会触发panic
	//关闭该ticker可以释放相关资源
	ticker := time.NewTicker(2*time.Second)
	go func(t *time.Ticker) {
		defer wg.Done()
		for {
			select {
				case received := <- t.C	:		//注意这里的返回值是时间类型
					fmt.Printf("get ticker, received:%s, current time:%s \n", received.Format(TEMPLATE), time.Now().Format(TEMPLATE))
					processOtherThing()
			}
		}
	}(ticker)

	wg.Wait()
}

//处理其他比较费时的任务
func processOtherThing() {
	fmt.Println("current goroutine is procesing other thing")
	time.Sleep(2100 * time.Millisecond)
}



