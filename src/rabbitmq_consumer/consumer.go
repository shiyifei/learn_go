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
		mq.Run()
	}()

	go func() {
		defer wg.Done()
		var mq1 *service.RabbitMQ
		mq1 = mq1.New("exchange_na", "topic", true)
		consumer1 := service.NewConsumer("simple:queue", "simple:#", true)
		mq1.RegisterReceiver(consumer1)
		mq1.Run()
	}()

	go func() {
		defer wg.Done()
		var mq2 *service.RabbitMQ
		mq2 = mq2.New("exchange.fanout", "fanout", true)
		consumer2 := service.NewConsumer("queue.fanout", "key.fanout", true)
		mq2.RegisterReceiver(consumer2)
		consumer21 := service.NewConsumer("queue:fanout", "key.fanout", true)
		mq2.RegisterReceiver(consumer21)
		mq2.Run()
	}()

	go func() {
		defer wg.Done()
		var mq3 *service.RabbitMQ
		mq3 = mq3.New("", "direct", false)
		consumerB := service.NewConsumer("queue.direct", "#.direct", false)
		mq3.RegisterReceiver(consumerB)
		mq3.Run()
	}()

	wg.Wait()


}
