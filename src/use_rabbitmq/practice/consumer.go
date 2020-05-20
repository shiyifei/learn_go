package practice

import "log"

func Consume() {
	conn, err := RabbitMQConn()
	ErrorHandling(err, "Failed to connect to RabbitMQ")

	defer conn.Close()

	ch, err := conn.Channel()
	ErrorHandling(err ,"failed to open a channel")

	defer ch.Close()

	q, err := ch.QueueDeclare("simple:queue", false, false, false, false, nil,)
	ErrorHandling(err, "failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
		)
	ErrorHandling(err, "Failed to register a consumer")

	go func() {
		for d:= range msgs {
			log.Printf("Received a message:%s \n", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages...")
	select {}

}