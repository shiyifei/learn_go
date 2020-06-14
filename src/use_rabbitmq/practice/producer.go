package practice

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
	"math/rand"
	"strconv"
	"sync"
	"time"
)


type simpleDemo struct {
	Name string `json:"name"`
	Addr string `json:"addr"`
	Age  int 	`json:"age"`
}

type target struct {
	queue string
	exchange string
	routingKey string
	durable bool
}

var wg sync.WaitGroup

func Produce() {
	wg.Add(4)
	go produceMsg("simple:queue", "exchange_na", true, "simple:queue")
	go produceMsg("queue.fanout", "exchange.fanout", true, "key.fanout")
	go produceMsg("queue.direct", "", true, "queue.direct")
	go produceMsg("jcque", "exchange_jc", true, "key:jc")
	wg.Wait()
}


func produceMsg(queue, exchange string, durable bool, routingKey string) {
	defer wg.Done()
	conn, err := RabbitMQConn()
	FailOnError(err, "Failed to connect to RabbitMQ")

	//关闭连接
	defer conn.Close()

	//当前连接打开一个Channel
	ch, err := conn.Channel()
	FailOnError(err ,"Failed to open a channel")

	//关闭通道
	defer ch.Close()



	//声明或创建一个队列用来保存消息
	q, err := ch.QueueDeclare(
		queue,			//queue name
		durable,		//durable
		false,			//delete when unused
		false,			//exclusive 独有的，排外的
		false,			//no-wait
		nil,			//arguments
		)
	FailOnError(err, "Failed to declare a queue")

	//构造发送的数据
	rand.Seed(time.Now().UnixNano())
	data := simpleDemo{
		Name : "user"+ strconv.Itoa(rand.Int()%1000),
		Addr : "Beijing",
		Age	 :  rand.Int()%50,
	}
	//数据转为json串
	dataBytes, err := json.Marshal(data)
	FailOnError(err, "struct to json failed")

	options := amqp.Publishing{
		ContentType: "text/plain",
		Body: dataBytes,
	}

	if durable {
		options.DeliveryMode = amqp.Persistent	//如果想让消息持久化，需要设置DeliveryMode为amqp.Persistent
	}
	//发布消息
	err = ch.Publish(
		exchange,		//exchange
		routingKey,	//routing key
		false,	//mandatory 是否是必选项目
		false,	//是否立即发送
		options,
	)

	log.Printf(" [x] Sent %s to queue:[%s]", dataBytes, q.Name)
	FailOnError(err, "Failed to publish a message")
}



func MultiSendMsg() {
	conn, err := RabbitMQConn()
	FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	targetArr := make([]target, 0)
	targetArr = append(targetArr, target{"simple:queue", "exchange_na", "simple:queue", true})
	targetArr = append(targetArr, target{"queue.fanout", "exchange.fanout", "key.fanout", true})
	targetArr = append(targetArr, target{"queue.direct", "", "queue.direct", false})
	targetArr = append(targetArr, target{"jcque", "exchange_jc", "key:jc", true})

	var wg sync.WaitGroup

	errList := make(chan error, 2 * len(targetArr))
	for _, t := range targetArr {
		wg.Add(1)
		go func(t target) {
			defer wg.Done()
			args := make(map[string]interface{})
			args["x-dead-letter-exchange"] = "exchange.fail"
			if t.queue == "jcque" {
				args["x-dead-letter-routing-key"] = "from.queue.jc"
			} else {
				args["x-dead-letter-routing-key"] = "from."+t.queue
			}

			//声明或创建一个队列用来保存消息
			q, err := ch.QueueDeclare(
				t.queue,			//queue name
				t.durable,		//durable
				false,			//delete when unused
				false,			//exclusive 独有的，排外的
				false,			//no-wait
				args,			//arguments
			)
			if err != nil {
				errList <- err
				return
			}
			//构造发送的数据
			rand.Seed(time.Now().UnixNano())

			for i:=0; i<10;i++ {
				data := simpleDemo{
					Name : "user"+ strconv.Itoa(rand.Int()%1000),
					Addr : "Beijing",
					Age	 :  rand.Int()%50,
				}
				//数据转为json串
				dataBytes, err := json.Marshal(data)
				FailOnError(err, "struct to json failed")

				options := amqp.Publishing{
					ContentType: "text/plain",
					Body: dataBytes,
				}

				if t.durable {
					options.DeliveryMode = amqp.Persistent	//如果想让消息持久化，需要设置DeliveryMode为amqp.Persistent
				}

				//发布消息
				err = ch.Publish(
					t.exchange,		//exchange
					t.routingKey,	//routing key
					false,	//mandatory 是否是必选项目
					false,	//是否立即发送
					options,
				)
				if err != nil {
					errList <- err
					return
				}
				log.Printf(" [x] Sent %s to queue:[%s],routerkey:%s", dataBytes, q.Name, t.routingKey)
			}

		}(t)
	}
	wg.Wait()
	close(errList)
	if len(errList) > 0{
		for err := range errList {
			FailOnError(err, "Failed send message")
		}
	}


}