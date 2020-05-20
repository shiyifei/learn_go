package practice

import "log"

func ProcessTask() {
	conn, err := RabbitMQConn()
	ErrorHandling(err, "Failed to connect to RabbitMQ")

	defer conn.Close()

	ch, err := conn.Channel()
	ErrorHandling(err ,"failed to open a channel")

	defer ch.Close()

	q, err := ch.QueueDeclare(
		"task:queue",
		false,
		false,
		false,
		false,
		nil,
	)
	ErrorHandling(err, "Failed to declare a queue")

	err = ch.Qos(
		1,
		0,
		false)
	ErrorHandling(err, "failed to set Qos")


	msgs, err := ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	ErrorHandling(err, "Failed to register a consumer")

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf("Received a message:%s", d.Body)
			log.Printf("Done")
			d.Ack(false)
		}
	}()
	log.Printf(" [*] Waiting for messages. \n")
	<-forever



}
