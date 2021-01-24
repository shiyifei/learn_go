/**
 * @Author:shiyf
 * @Date: 2021/1/24 22:58
 **/

package practice

import (
	"fmt"
	"github.com/Shopify/sarama"
)

func SendMessage() {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Partitioner = sarama.NewRandomPartitioner

	client, err := sarama.NewClient([]string{"192.168.1.9:9092"}, config)
	defer client.Close()

	if err != nil {
		panic(err)
	}

	producer, err := sarama.NewSyncProducerFromClient(client)
	if err != nil {
		panic(err)
	}

	partition, offset, err :=
		producer.SendMessage(&sarama.ProducerMessage{Topic:"test", Key:nil, Value:sarama.StringEncoder("hello,world")})
	if err != nil {
		panic(err)
	}
	fmt.Println("partition:",partition)
	fmt.Println("offset", offset)
}


func ReceiveMsg() {

}