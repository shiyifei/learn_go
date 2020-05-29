package practice

/**
	定义一个任务的生产者,用于生产任务消息
 */

import (
	"fmt"
	"log"
	"strings"
	"github.com/streadway/amqp"
)


func bodyFrom(args []string) string {
	var s string
	if len(args) <= 0 {
		s = "no task"
	} else {
		s = strings.Join(args, " ")
	}
	return s
}

//生成任务消息
func GenerateTask(args []string) {
	fmt.Println("arrive in GenerateTask()")
	conn, err := RabbitMQConn()
	FailOnError(err, "Failed to connect to RabbitMQ")

	defer conn.Close()

	ch, err := conn.Channel()
	FailOnError(err ,"failed to open a channel")

	defer ch.Close()

	q, err := ch.QueueDeclare(
		"task:queue",
		false,
		false,
		false,
		false,
		nil,
	)
	FailOnError(err, "Failed to declare a queue")

	body := bodyFrom(args)
	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType:"text/plain",
			DeliveryMode: amqp.Persistent,
			Body : []byte(body),
		})
	FailOnError(err, "Failed to generate a task")
	log.Printf("sent %s \n", body)
}

