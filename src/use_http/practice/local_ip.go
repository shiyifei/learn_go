/**
 * 获取本机的IP地址
 */
package practice

import (
	"net"
	"fmt"
	"strings"
)

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