package practice

import (
	"log"
	"net/http"
	"time"
)

const (
	numPollers = 2						//表示最终开启几个协程去模拟测试网站服务操作
	pollInterval = 60 * time.Second		//允许多长时间间隔才能再次发起请求
	statusInterval = 10 * time.Second	//表示多长时间间隔去输出每个站点的响应状态
	errTimeout = 10 * time.Second		//表示当出错时的超时回退时间
)

//表示参与测试服务状态的站点
var urls = []string{
	"https://www.baidu.com",
	"https://www.qq.com/",
	"http://192.168.56.102/",
}

//该结构体表示一个站点最终的响应状态
type State struct {
	url string
	status string
}

//创建定时器，监控每一个请求的处理结果
//注意这里的State实际上是一个记录请求状态的只允许单向写入的通道变量，也是该方法最终生成的结果
//这里的State实际上先执行了初始化通道的操作，后来有其他协程向通道中写入了值，本方法中的goroutine又读取了goroutine中的值
func StateMonitor(updateInterval time.Duration) chan<- State {
	updates := make(chan State)
	urlStatus := make(map[string]string)
	ticker := time.NewTicker(updateInterval)
	go func() {
		for {
			select {
				case <-ticker.C:
					logState(urlStatus)
				case s := <-updates:
						log.Printf("read from channle State, s:%+v \n", s)
						urlStatus[s.url] = s.status

			}
		}
	}()
	return updates
}

func logState(s map[string]string) {
	log.Println("Current state:[")
	for k, v := range s {
		log.Printf("%s %s", k, v)
	}
	log.Println("]")
}

type Resource struct {
	url string
	errCount int
}

func (r *Resource) Poll() string {
	resp, err := http.Head(r.url)
	if err != nil {
		log.Println("Error", r.url, err)
		r.errCount++
		return err.Error()
	}
	r.errCount = 0
	return resp.Status
}

//参数变量是一个只允许写入值的通道类型
//针对某一个Resource 休眠一段时间 (出错次数越多休眠时间越长)
func (r *Resource) Sleep(done chan <- *Resource) {
	log.Printf("arrive in Sleep()，r:%+v \n", *r)
	time.Sleep(pollInterval + errTimeout * time.Duration(r.errCount) )
	done <- r
}

/**
in是一个只允许单向读取的Resource类型通道，out 是一个只允许单向写入的Resource类型通道
status 是一个只允许单向写入的State类型通道
本方法针对输入的站点列表开始测试站点的响应情况，记录响应状态。
 */
func Poller( in <-chan *Resource, out chan<- *Resource, status chan<- State) {
	for r := range in {
		s := r.Poll()
		status <- State{r.url, s}
		out <- r
	}
}

//演示通过通信的方式来完成共享内存的操作实例
func ShareMemoryByCommunication() {
	pending, complete := make(chan *Resource), make(chan *Resource)

	//初始化状态通道，创建定时器，监控每一个请求的处理结果
	status := StateMonitor(statusInterval)

	//开启两个协程，分别测试每一个站点的响应情况
	for i:=0;i<numPollers;i++ {
		go Poller(pending, complete, status)
	}

	//向待处理的资源通道中添加起始站点列表
	go func() {
		for _, url := range urls {
			pending <- &Resource{url:url}
		}
	}()

	//最终输出结果
	for r := range complete {
		go r.Sleep(pending)
	}
}
