package practice

import(
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

func RabbitMQConn() (conn *amqp.Connection, err error) {
	//RabbitMQ分配的用户名称
	var user string = "admin"
	var pwd string = "manager"
	var host string = "192.168.56.102"
	var port string = "5672"

	url := fmt.Sprintf("amqp://%s:%s@%s:%s/", user, pwd, host, port)
	conn, err = amqp.Dial(url)
	return
}

func ErrorHandling(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}