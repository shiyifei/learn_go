package service

import(
"fmt"
"github.com/streadway/amqp"
"log"
	"net"
	"strings"
)

func RabbitMQConn() (conn *amqp.Connection, err error) {
	//RabbitMQ分配的用户名
	var user string = "admin"
	//RabbitMQ分配的密码
	var pwd string = "manager"

	//RabbitMQ Broker的ip地址
	var host string = "192.168.56.110"

	localIp := GetLocalIp()
	if localIp == "192.168.1.102" {
		host = "192.168.1.102"
	}

	//RabbitMQ Broker监听的端口
	var port string = "5673"

	url := fmt.Sprintf("amqp://%s:%s@%s:%s/", user, pwd, host, port)

	//新建一个连接
	conn, err = amqp.Dial(url)
	return
}

/**
	输出异常信息
 */
func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func GetLocalIp() string {
	conn, err := net.Dial("udp", "www.baidu.com:80")
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().String()

	localIp := strings.Split(localAddr, ":")[0]
	fmt.Println("localAddr:", localAddr, "localIP:",localIp)
	return localIp
}
