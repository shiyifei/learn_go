package practice

import(
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/streadway/amqp"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"
)

type JsonData struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
}

func RabbitMQConn() (conn *amqp.Connection, err error) {
	//RabbitMQ分配的用户名
	var user string = "admin"
	//RabbitMQ分配的密码
	var pwd string = "manager"

	//RabbitMQ Broker的ip地址
	var host string = "192.168.1.102"

	localIp := GetLocalIp()
	if localIp == "192.168.56.107" {
		host = "192.168.56.110"
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
		log.Printf("%s: %s \n", msg, err)
	}
}

func SendPostRequest(url, message string) (bool, error){
	var jsonStr = []byte(message)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	// req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	//fmt.Println("response Status:", resp.Status)
	//fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	var data JsonData
	json.Unmarshal(body, &data)
	if data.Code == 1 {
		return true, nil
	} else {
		return false, errors.New(data.Msg)
	}
}



func GetLocalIp() string {
	/*var ips = make(map[string]string)
	ips, _ = Ips()
	fmt.Printf("ips:%+v \n", ips)

	addrs, err := net.InterfaceAddrs()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, address := range addrs {
		fmt.Printf("address:%+v \n", address)
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			fmt.Printf("ipnet:%+v \n", ipnet)
			if ipnet.IP.To4() != nil {
				fmt.Println(ipnet.IP.String())
			}

		}
	}*/

	conn, err := net.Dial("udp", "www.baidu.com:80")
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().String()

	localIp := strings.Split(localAddr, ":")[0]
	//fmt.Println("localAddr:", localAddr, "localIP:",localIp)
	return localIp
}

//获取全部网卡的全部IP
func Ips() (map[string]string, error) {

	ips := make(map[string]string)

	//返回 interface 结构体对象的列表，包含了全部网卡信息
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	//遍历全部网卡
	for _, i := range interfaces {

		// Addrs() 方法返回一个网卡上全部的IP列表
		address, err := i.Addrs()
		if err != nil {
			return nil, err
		}

		//遍历一个网卡上全部的IP列表，组合为一个字符串，放入对应网卡名称的map中
		for _, v := range address {
			ips[i.Name] += v.String() + " "
		}
	}
	return ips, nil
}