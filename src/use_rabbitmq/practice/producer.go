package practice

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
)


type simpleDemo struct {
	Name string `json:"name"`
	Addr string `json:"addr"`
}


func Produce() {
	conn, err := RabbitMQConn()
	ErrorHandling(err, "Failed to connect to RabbitMQ")

	defer conn.Close()

	ch, err := conn.Channel()
	ErrorHandling(err ,"Failed to open a channel")

	defer ch.Close()

	q, err := ch.QueueDeclare(
		"simple:queue",
		false,
		false,
		false,
		false,
		nil,
		)
	ErrorHandling(err, "Failed to declare a queue")

	data := simpleDemo{
		Name : "wangyueyang",
		Addr : "Beijing",
	}

	dataBytes, err := json.Marshal(data)
	ErrorHandling(err, "struct to json failed")

	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body: dataBytes,
		})
	log.Printf(" [x] Sent %s", dataBytes)
	ErrorHandling(err, "Failed to publish a message")
}