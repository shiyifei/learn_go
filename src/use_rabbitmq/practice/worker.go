package practice

import "log"

func ProcessTask() {
	conn, err := RabbitMQConn()
	FailOnError(err, "Failed to connect to RabbitMQ")

	defer conn.Close()

	ch, err := conn.Channel()
	FailOnError(err ,"failed to open a channel")

	defer ch.Close()

	q, err := ch.QueueDeclare(
		"queue.direct",
		false,
		false,
		false,
		false,
		nil,
	)
	FailOnError(err, "Failed to declare a queue")

	//函数原型：func (ch *Channel) Qos(prefetchCount, prefetchSize int, global bool) error
	// rabbitMQ提供了一种qos（服务质量保证）的功能，
	// 即非自动确认消息的前提下，如果有一定数目的消息（通过consumer或者Channel设置qos）未被确认，不进行新的消费。
	err = ch.Qos(
		10,	//设置为N，告诉rabbitMQ不要同时给一个消费者推送多于N个消息，即一旦有N个消息还没有ack，则consumer将block掉，直到有消息ack
		0,
		false)
	FailOnError(err, "failed to set Qos, queue name:"+q.Name)


	msgs, err := ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	FailOnError(err, "Failed to register a consumer")

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf("Received a message:%s from queue:[%s]", d.Body, q.Name)
			log.Printf("Done")
			d.Ack(false)
		}
	}()
	log.Printf(" [*] Waiting for messages. \n")
	<-forever



}
