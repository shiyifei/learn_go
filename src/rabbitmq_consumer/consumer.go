package main

import (
	"rabbitmq_consumer/service"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	wg.Add(4)

	go func() {
		defer wg.Done()

		var mq *service.RabbitMQ
		mq = mq.New("exchange_jc", "topic", true)
		consumerA := service.NewConsumer("jcque", "key:**", true)
		mq.RegisterReceiver(consumerA)
		consumerB := service.NewConsumer("jcque", "jc#.xin.com", true)
		mq.RegisterReceiver(consumerB)
		mq.Run()
	}()

	go func() {
		defer wg.Done()
		var mq *service.RabbitMQ
		mq = mq.New("exchange_na", "topic", true)
		consumerA := service.NewConsumer("simple:queue", "simple:#", true)
		mq.RegisterReceiver(consumerA)
		mq.Run()
	}()

	go func() {
		defer wg.Done()
		var mq *service.RabbitMQ
		mq = mq.New("exchange.fanout", "fanout", true)
		consumerA := service.NewConsumer("queue:fanout", "", true)
		mq.RegisterReceiver(consumerA)
		mq.Run()
	}()

	go func() {
		defer wg.Done()

		var mq *service.RabbitMQ
		mq = mq.New("", "direct", false)
		consumerB := service.NewConsumer("queue.direct", "#.direct", false)
		mq.RegisterReceiver(consumerB)
		mq.Run()
	}()

	wg.Wait()


}
