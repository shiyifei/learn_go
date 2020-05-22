package practice

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
	"math/rand"
	"strconv"
	"time"
)


type simpleDemo struct {
	Name string `json:"name"`
	Addr string `json:"addr"`
	Age  int 	`json:"age"`
}

func Produce() {
	produceMsg("simple:queue", "exchange_na", true, "simple:queue")
	//produceMsg("queue_na", "change_na", true)
	produceMsg("jcque", "exchange_jc", true, "key:jc")
	produceMsg("jcque", "exchange_jc", true, "key:jc")
	produceMsg("jcque", "exchange_jc", true, "key:jc")
}


func produceMsg(queue, exchange string, durable bool, routingKey string) {
	conn, err := RabbitMQConn()
	ErrorHandling(err, "Failed to connect to RabbitMQ")

	//关闭连接
	defer conn.Close()

	//当前连接打开一个Channel
	ch, err := conn.Channel()
	ErrorHandling(err ,"Failed to open a channel")

	//关闭通道
	defer ch.Close()

	//声明或创建一个队列用来保存消息
	q, err := ch.QueueDeclare(
		queue,	//queue name
		durable,			//durable
		false,			//delete when unused
		false,			//exclusive 独有的，排外的
		false,			//no-wait
		nil,			//arguments
		)
	ErrorHandling(err, "Failed to declare a queue")
	rand.Seed(time.Now().UnixNano())
	data := simpleDemo{
		Name : "user"+ strconv.Itoa(rand.Int()%1000),
		Addr : "Beijing",
		Age	 :  rand.Int()%50,

	}
	//数据转为json串
	dataBytes, err := json.Marshal(data)
	ErrorHandling(err, "struct to json failed")

	//发布消息
	err = ch.Publish(
		exchange,		//exchange
		routingKey,	//routing key
		false,	//mandatory 是否是必选项目
		false,	//是否立即发送
		amqp.Publishing{
			ContentType: "text/plain",
			Body: dataBytes,
		})
	log.Printf(" [x] Sent %s to queue:[%s]", dataBytes, q.Name)
	ErrorHandling(err, "Failed to publish a message")
}