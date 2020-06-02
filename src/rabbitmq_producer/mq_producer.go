package main

import (
	"fmt"
	"rabbitmq_producer/service"
)

func main() {

	producer := service.NewProducer("queue:new","amqp://admin:manager@192.168.56.110:5673/")
	producer.Connect("amqp://admin:manager@192.168.56.110:5673/")
	//fmt.Printf("producer:%+v \n", producer)
	str := `{"name":"user483","addr":"Beijing","age":30}`
	producer.Push([]byte(str))
	fmt.Println("after producer.Push()")
}
