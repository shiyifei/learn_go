package practice

import (
	"fmt"
	"log"
	"sync"
	"time"

	//"time"
)

type Source struct {
	queue string
	bindingKey string
	exchange string
	durable bool
}

var wgC sync.WaitGroup

func MultiConsume() {
	conn, err := RabbitMQConn()
	FailOnError(err, "Failed to connect to RabbitMQ")

	defer conn.Close()

	ch, err := conn.Channel()
	FailOnError(err ,"failed to open a channel")

	defer ch.Close()


	//函数原型：func (ch *Channel) Qos(prefetchCount, prefetchSize int, global bool) error
	// rabbitMQ提供了一种qos（服务质量保证）的功能，
	// 即非自动确认消息的前提下，如果有一定数目的消息（通过consumer或者Channel设置qos）未被确认，不进行新的消费。
	err = ch.Qos(
		1, //设置为N，告诉rabbitMQ不要同时给一个消费者推送多于N个消息，即一旦有N个消息还没有ack，则consumer将block掉，直到有消息ack
		0,
		true)

	FailOnError(err, "failed to set Qos")

	sourceArr := make([]Source, 0)
	sourceArr = append(sourceArr, Source{"simple:queue", "simple:#", "exchange_na", true})
	sourceArr = append(sourceArr, Source{"jcque", "key:**", "exchange_jc", true})

	//fanout类型的exchange不需要绑定bindingKey
	sourceArr = append(sourceArr, Source{"queue.fanout", "", "exchange.fanout", true})
	sourceArr = append(sourceArr, Source{"queue.direct", "#.direct", "", false})

	/*	err = ch.ExchangeDeclare(exchange, "topic", true, false, false, false, nil)
		if err != nil {
			log.Fatalf("exchange.declare: %s", err)
		}*/


	forever := make(chan bool)

	var wg sync.WaitGroup
	wg.Add(len(sourceArr))
	for _, s := range sourceArr {
		go func(s Source) {
			defer wg.Done()
			q, err := ch.QueueDeclare(s.queue, s.durable, false, false, false, nil,)
			FailOnError(err, "failed to declare a queue")

			//加上这一段之后，有时候会报错，先去掉
			/*if s.exchange != "" {
				err = ch.QueueBind(s.queue, s.bindingKey, s.exchange, false, nil)
				FailOnError(err, "failed to bind a queue")
			}*/

			//定义一个消费者
			msgs, err := ch.Consume(
				q.Name,	//queue name
				"",		//consumer
				false,	//auto-ack
				false,	//exclusive
				false,	//no-local
				false,	//no-wait
				nil,	//arguments
			)
			FailOnError(err, "Failed to register a consumer,queue name:"+q.Name)

			go func() {
				for d := range msgs {
					log.Printf("Received a message %s from queue:[%s] \n", d.Body, q.Name)
					//time.Sleep(200*time.Millisecond)
					//先给出应答，再调用接口发送消息，保证不影响后续并发
					d.Ack(false)
					var wg1 sync.WaitGroup
					wg1.Add(1)
					go callBack(s.exchange, s.queue, s.bindingKey, string(d.Body), &wg1)
					wg1.Wait()
					//writeToDB(s.exchange, s.queue, s.bindingKey, string(d.Body))
					//callBackA(s.exchange, s.queue, s.bindingKey, string(d.Body))
				}
			}()

			log.Printf(" queue:[%s] Waiting for messages...", q.Name)
		}(s)
	}
	wg.Wait()


	<-forever
}

func callBack(exchange, queue, bindingKey, message string, wg1 *sync.WaitGroup) {
	defer wg1.Done()
	_, err := SendPostRequest("http://192.168.56.107:8100/mq/consume", message)
	if err != nil {
		FailOnError(err, "send post request error one time")
		writeToDB(exchange, queue, bindingKey, message)

		time.Sleep(200*time.Millisecond)
		//失败重试一次
		_, err = SendPostRequest("http://192.168.56.107:8100/mq/consume", message)
		if err != nil {
			FailOnError(err, "send post request error, two times")

			//处理失败或调用接口失败时，写入数据库
			writeToDB(exchange, queue, bindingKey, message)
		}
	}
}

/**
CREATE TABLE `failed_message` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `exchange` varchar(30) NOT NULL DEFAULT '' COMMENT 'exchange name',
  `queue` varchar(30) NOT NULL DEFAULT '' COMMENT 'queue name',
  `binding_key` varchar(30) NOT NULL DEFAULT '' COMMENT 'binding key',
  `message` varchar(512) NOT NULL DEFAULT '' COMMENT 'message',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
 */

func writeToDB(exchange,queue,bindingKey,message string) {
	//fmt.Println("in writeToDB(),111")
	stmt, err := SqlDB.Prepare("insert into failed_message(exchange,queue,binding_key,message) values(?,?,?,?)")
	FailOnError(err, "prepare error")
	res,err := stmt.Exec(exchange, queue, bindingKey, message)
	FailOnError(err, "insert error")
	fmt.Println("res,",res)
}



func Consume() {
	fmt.Println("arrive in Consume(),111")

	wgC.Add(8)
	go consumeBind("simple:queue", "simple:#", "exchange_na", true)
	go consumeBind("jcque", "key:**", "exchange_jc", true)
	go consumeBind("queue.fanout", "#.fanout", "exchange.fanout", true)
	go consumeBind("queue.direct", "queue.direct", "", false)
	wgC.Wait()

	forever := make(chan bool)
	<-forever

}

func consumeBind(queue, routingKey, exchange string, durable bool) {
	defer wgC.Done()
	fmt.Println("arrive in ConsumeBind(),111,",queue)
	conn, err := RabbitMQConn()
	FailOnError(err, "Failed to connect to RabbitMQ")

	fmt.Println("arrive in ConsumeBind(),222")

	defer conn.Close()

	ch, err := conn.Channel()
	FailOnError(err ,"failed to open a channel")

	defer ch.Close()

	fmt.Println("arrive in ConsumeBind(),333")

/*	err = ch.ExchangeDeclare(exchange, "topic", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("exchange.declare: %s", err)
	}*/

	fmt.Println("arrive in ConsumeBind(),444")

	q, err := ch.QueueDeclare(queue, durable, false, false, false, nil,)
	FailOnError(err, "failed to declare a queue")

	fmt.Println("arrive in ConsumeBind(),555")

	if exchange != "" {
		err = ch.QueueBind(queue, routingKey, exchange, false, nil)
		FailOnError(err, "failed to bind a queue")
		fmt.Println("arrive in ConsumeBind(),666")
	}



	//定义一个消费者
	msgs, err := ch.Consume(
		q.Name,	//queue name
		"",		//consumer
		true,	//auto-ack
		false,	//exclusive
		false,	//no-local
		false,	//no-wait
		nil,	//arguments
	)
	FailOnError(err, "Failed to register a consumer")

	fmt.Println("arrive in ConsumeBind(),777")
	go func() {
		defer wgC.Done()
		for d:= range msgs {
			//fmt.Println("arrive in ConsumeBind(),888")
			log.Printf("Received a message %s from queue:[%s] \n", d.Body, q.Name)
		}
	}()

	log.Printf(" queue[%s] Waiting for messages...", q.Name)
}