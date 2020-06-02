package config

import (
	"fmt"
	"net"
	"strings"
)

var ConnectStr string

func init() {
		//RabbitMQ分配的用户名
		var user string = "admin"
		//RabbitMQ分配的密码
		var pwd string = "manager"

		//RabbitMQ Broker的ip地址
		var host string = "192.168.56.110"

		/*localIp := GetLocalIp()
		if localIp == "192.168.1.102" {
			host = "192.168.1.102"
		}*/

		//RabbitMQ Broker监听的端口
		var port string = "5673"

		ConnectStr = fmt.Sprintf("amqp://%s:%s@%s:%s/", user, pwd, host, port)
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

