package service

import (
	"log"
)

type Receiver interface {
	QueueName() string
	BindingKey() string
	Durable() bool
	OnError(error)
	OnReceive(string,[]byte)bool
}

type Consumer struct {
	queueName string
	bindingKey string
	durable bool
}

func NewConsumer(queueName, bindingKey string, durable bool) Consumer {
	var recv Consumer
	recv.queueName = queueName
	recv.bindingKey = bindingKey
	recv.durable = durable
	return recv
}

func (c Consumer) QueueName() string {
	return c.queueName
}

func (c Consumer) BindingKey() string {
	return c.bindingKey
}

func (c Consumer) Durable() bool {
	return c.durable
}

func (c Consumer) OnError(err error) {
	log.Println(err.Error())
}

func (c Consumer) OnReceive(exchange string, message []byte) bool {
	log.Printf("receive message:%s, exchange:%s,queue:%s,binding key:%s \n", message, exchange, c.queueName, c.bindingKey)
	return true
}



