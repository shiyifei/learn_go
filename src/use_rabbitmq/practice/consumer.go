package practice

import (
	"fmt"
	"log"
)



/*func Consume() {
	conn, err := RabbitMQConn()
	ErrorHandling(err, "Failed to connect to RabbitMQ")

	defer conn.Close()

	ch, err := conn.Channel()
	ErrorHandling(err ,"failed to open a channel")

	defer ch.Close()

	q, err := ch.QueueDeclare("simple:queue", false, false, false, false, nil,)
	ErrorHandling(err, "failed to declare a queue")

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
	ErrorHandling(err, "Failed to register a consumer")

	go func() {
		for d := range msgs {
			log.Printf("Received a message:%s from %s \n", d.Body, q.Name)
		}
	}()

	log.Printf(" [*] Waiting for messages...")
	select {}
}*/

func Consume() {
	fmt.Println("arrive in Consume(),111")

	consumeBind("simple:queue", "simple:queue", "exchange_na")
	consumeBind("jcque", "key:jc", "exchange_jc")


}

func consumeBind(queue, routingKey, exchange string) {
	fmt.Println("arrive in ConsumeBind(),111,",queue)
	conn, err := RabbitMQConn()
	ErrorHandling(err, "Failed to connect to RabbitMQ")

	fmt.Println("arrive in ConsumeBind(),222")

	defer conn.Close()

	ch, err := conn.Channel()
	ErrorHandling(err ,"failed to open a channel")

	defer ch.Close()

	fmt.Println("arrive in ConsumeBind(),333")

/*	err = ch.ExchangeDeclare(exchange, "topic", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("exchange.declare: %s", err)
	}*/

	fmt.Println("arrive in ConsumeBind(),444")

	q, err := ch.QueueDeclare(queue, true, false, false, false, nil,)
	ErrorHandling(err, "failed to declare a queue")

	fmt.Println("arrive in ConsumeBind(),555")

	err = ch.QueueBind(queue, routingKey, exchange, false, nil)
	ErrorHandling(err, "failed to bind a queue")

	fmt.Println("arrive in ConsumeBind(),666")


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
	ErrorHandling(err, "Failed to register a consumer")

	fmt.Println("arrive in ConsumeBind(),777")

	go func() {
		for d:= range msgs {
			fmt.Println("arrive in ConsumeBind(),888")
			log.Printf("Received a message [%s] from queue:%s \n", d.Body, q.Name)
		}
	}()

	log.Printf(" [*] Waiting for messages...")
	//select {}

}