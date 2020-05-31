package service

import (
	"sync"
	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	wg sync.WaitGroup
	channel *amqp.Channel
	exchangeName string
	exchangeType string
	durable bool
	receivers []Receiver
}


func (mq *RabbitMQ) New(exchangeName, exchangeType string, durable bool) *RabbitMQ {
	return &RabbitMQ{
		exchangeName:exchangeName,
		exchangeType:exchangeType,
		durable:durable,
	}
}

func (mq *RabbitMQ) prepareExchange() error {
	//声明exchange
	err := mq.Channel.ExchangeDeclare(
			mq.exchangeName,	//exchange
			mq.exchangeType,	//type
			mq.durable,			//durable
			false,				//auto delete
			false,				//internal
			false,				//nowait
			nil,				//args
		)
	if err != nil {
		return err
	}
	return nil
}

func (mq *RabbitMQ) run() {

}
