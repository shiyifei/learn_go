package practice

import(
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

func RabbitMQConn() (conn *amqp.Connection, err error) {
	//RabbitMQ分配的用户名
	var user string = "admin"
	//RabbitMQ分配的密码
	var pwd string = "manager"

	//RabbitMQ Broker的ip地址
	var host string = "192.168.56.102"

	//RabbitMQ Broker监听的端口
	var port string = "5672"

	url := fmt.Sprintf("amqp://%s:%s@%s:%s/", user, pwd, host, port)

	//新建一个连接
	conn, err = amqp.Dial(url)
	return
}

/**
	输出异常信息
 */
func ErrorHandling(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}