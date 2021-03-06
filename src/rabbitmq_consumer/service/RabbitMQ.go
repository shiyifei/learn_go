package service

import (
	"errors"
	"fmt"
	"rabbitmq_consumer/config"
	"sync"
	"github.com/streadway/amqp"
	"log"
	"time"
)
var mutex sync.Mutex

type RabbitMQ struct {
	wg sync.WaitGroup
	connection *amqp.Connection
	channel *amqp.Channel
	notifyClose chan *amqp.Error
	isConnected bool
	exchangeName string
	exchangeType string
	durable bool
	receivers []Receiver
}

const (
	reconnectTimes = 180			//连接断开后可以重连多少次
	reconnectDelay = 1 				//单位：time.Second 连接断开后多久重连
	resendDelay = 1* time.Second	//消息发送失败后多久重发
	resendTimes = 1				//消息重发次数
	dateTemplate = "2006-01-02 15:04:05.000"  //go时间格式 加入毫秒显示 如果是6个0则是微秒，3个0表示不足三位时左侧会补齐0
)


func (mq *RabbitMQ) New(exchangeName, exchangeType string, durable bool) *RabbitMQ {
	client := RabbitMQ{
		exchangeName:exchangeName,
		exchangeType:exchangeType,
		durable:durable,
	}
	client.notifyClose = make(chan *amqp.Error)
	go client.handleReConnect(config.ConnectStr)
	return &client
}

func (mq *RabbitMQ) handleReConnect(addr string) {
	var disconnect *amqp.Error

	for {
		select {
			case disconnect = <-mq.notifyClose:
				log.Println("fail to connect rabbitmq server")
				mq.connection = nil
				mq.channel = nil
				mq.reconnect(addr, disconnect)
		}
	}
}

func (mq *RabbitMQ) reconnect(addr string, disconnect *amqp.Error) {
	var retryTimes int = 0
	log.Println("Attampting to connect:",addr,"disconnect:",disconnect)
	for !mq.connect(addr) {
		retryTimes += 1
		if retryTimes >= reconnectTimes {
			return
		}
		log.Println("Failed to connect. Retrying... retryTimes:",retryTimes)

		duration := time.Duration(retryTimes * reconnectDelay) * time.Second
		time.Sleep(duration)
	}
	if disconnect != nil && mq.connection != nil && !mq.connection.IsClosed() {
		mq.prepareExchange()
		if len(mq.receivers) > 0 {
			for _, receiver := range mq.receivers {
				mq.wg.Add(1)
				go mq.listen(receiver)
			}
			mq.wg.Wait()
		}
	}
}

func (mq *RabbitMQ) connect(addr string) bool {
	mutex.Lock()
	defer mutex.Unlock()
	conn, err := amqp.Dial(addr)
	if err != nil {
		return false
	}
	ch, err := conn.Channel()
	if err != nil {
		return false
	}

	mq.connection = conn
	mq.channel = ch
	mq.channel.NotifyClose(mq.notifyClose)		//加入监听事件

	fmt.Printf("arrive in connect(),receivers:%+v \n", mq.receivers)

	mq.isConnected = true
	return true
}

func (mq *RabbitMQ) prepareExchange() error {
	if mq.exchangeName == "" {
		return nil
	}
	//声明exchange
	err := mq.channel.ExchangeDeclare(
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

func (mq *RabbitMQ) Run() {
	log.Println("in Run(),exchange:", mq.exchangeName)
	defer mq.Close()
	if !mq.connect(config.ConnectStr) {
		log.Println("can not connect rabbitmq server")
		return
	}
	mq.prepareExchange()

	for {
		for _, receiver := range mq.receivers {
			mq.wg.Add(1)
			go mq.listen(receiver)
		}
		mq.wg.Wait()

		time.Sleep(300*time.Millisecond)
	}
}

func (mq *RabbitMQ) RegisterReceiver(receiver Receiver) {
	mq.receivers = append(mq.receivers, receiver)
}

func (mq *RabbitMQ) listen(receiver Receiver) {
	defer mq.wg.Done()

	//fmt.Printf("arrive in listen(), receiver:%+v \n", receiver)


	queueName := receiver.QueueName()
	routerKey := receiver.BindingKey()
	durable := receiver.Durable()

	if mq.channel == nil {
		return
	}

	_, err := mq.channel.QueueDeclare(
		queueName,	//queue name
		durable,	//durable
		false,		//delete when unused
		false,		//exclusive
		false,		//no-wait
		nil,		//argument
	)
	if err != nil {
		receiver.OnError(fmt.Errorf("初始化队列 %s 失败：%s", queueName, err.Error()))
	}

	//fmt.Println("routerKey:",routerKey)

	//fmt.Printf("arrive in listen(), after queueDeclare() \n")

	if mq.exchangeName != "" {
		err = mq.channel.QueueBind(
			queueName,
			routerKey,
			mq.exchangeName,
			false,
			nil,
		)
		if err != nil {
			receiver.OnError(errors.New(fmt.Sprintf("绑定队列 [%s-%s] 到交换机%s失败：%s", queueName, routerKey, mq.exchangeName,err.Error())))
		}
	}


	//fmt.Printf("arrive in listen(), after queueDeclar() \n")

	mq.channel.Qos(1, 0, true)	//确保rabbitmq会一个一个消费消息

	msgs, err := mq.channel.Consume(
		queueName,	//queue name
		"",			//consumer
		false,		//auto-ack
		false,		//exclusive
		false,		//no-local
		false,		//no-wait
		nil,		//args
		)

	//fmt.Printf("arrive in listen(), after Consume() \n")

	for msg := range msgs {
		//fmt.Printf("receive message:%s\n", msg.Body)
		msg.Ack(false)
		var wg1 sync.WaitGroup
		wg1.Add(1)
		receiver.OnReceive(mq.exchangeName, string(msg.Body), &wg1)
		wg1.Wait()
	}

}

func (mq *RabbitMQ) Close() error {
	if !mq.isConnected {
		return errors.New("already closed: not connected to rabbitmq server")
	}

	err := mq.channel.Close()
	if err != nil {
		return err
	}

	err = mq.connection.Close()
	if err != nil {
		return err
	}
	mq.isConnected = false
	return nil
}





