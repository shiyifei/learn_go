package common

import (
	"fmt"
	"log"
	"net"
	"strings"
)

/**
	抛出错误
 */
func CheckErr(err error) {
	if err != nil {
		log.Fatal(err.Error())
		panic(err)
	}
}

/**
	在最终main函数中捕获异常
 */
func TryRecover() {
	defer func() {
		r := recover()
		err, ok := r.(error)
		if ok {
			fmt.Println("捕获了这个错误！", err)
		} else {
			panic(r)
		}
	}()
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