package service

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type Receiver interface {
	QueueName() string
	BindingKey() string
	Durable() bool
	OnError(error)
	OnReceive(string,string, *sync.WaitGroup)bool
}

type Consumer struct {
	queueName string
	bindingKey string
	durable bool
}

func NewConsumer(queueName, bindingKey string, durable bool) Consumer {
	var recv Consumer
	recv.queueName = queueName
	recv.bindingKey = bindingKey
	recv.durable = durable
	return recv
}

func (c Consumer) QueueName() string {
	return c.queueName
}

func (c Consumer) BindingKey() string {
	return c.bindingKey
}

func (c Consumer) Durable() bool {
	return c.durable
}

func (c Consumer) OnError(err error) {
	log.Println(err.Error())
}

/*func (c Consumer) OnReceive(exchange string, message []byte) bool {
	log.Printf("receive message:%s, exchange:%s,queue:%s,binding key:%s \n", message, exchange, c.queueName, c.bindingKey)
	return true
}*/


func (c Consumer) OnReceive(exchange, message string, wg1 *sync.WaitGroup) bool {
	defer wg1.Done()
	_, err := SendPostRequest("http://192.168.56.106:8100/mq/consume", message)
	if err != nil {
		FailOnError(err, "send post request error once:")
		writeToDB(exchange, c.queueName, c.bindingKey, message, err.Error())

		ticker := time.NewTicker(resendDelay)
		defer ticker.Stop()

		var wg sync.WaitGroup
		wg.Add(1)
		var times int = 0

		go func(t *time.Ticker) {
			defer wg.Done()
			for {
				select {
				case received := <- t.C	:		//注意这里的返回值是时间类型
					fmt.Printf("get ticker, received:%s, current time:%s \n", received.Format(dateTemplate), time.Now().Format(dateTemplate))
					//失败重试两次
					_, err = SendPostRequest("http://192.168.56.106:8100/mq/consume", message)
					if err != nil {
						FailOnError(err, "send json request error")
					}
					times++
					if times >= resendTimes {
						return
					}
				}
			}
		}(ticker)
		wg.Wait()
		return false
	}
	return true
}

func writeToDB(exchange,queue,bindingKey,msg,errInfo string) {
	//fmt.Println("in writeToDB(),111")
	stmt, err := SqlDB.Prepare("insert into failed_message(exchange,queue,binding_key,message,error) values(?,?,?,?,?)")
	FailOnError(err, "prepare error")
	_,err = stmt.Exec(exchange, queue, bindingKey, msg, errInfo)
	FailOnError(err, "insert error")
}


