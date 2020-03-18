package practice

import (
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"net"
	"sync"
	"time"
)

const (
	SERVER_NETWORK = "tcp"
	SERVER_ADDRESS = "192.168.56.102:8085"
	DELIMITER = '\t'
	logSn = 666
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go clientDoWork(11, wg)
	wg.Wait()
}

func clientDoWork(id int, wg sync.WaitGroup) {
	conn, err := net.DialTimeout(SERVER_NETWORK, SERVER_ADDRESS, 5 * time.Second)
	if err != nil {
		printLog("Dial Error: %s (Client[%d]) \n", err, id)
		return
	}
	defer conn.Close()

	printLog("Connected to server. (remote Address:%s, local address:%s) (Client[%d]) \n", conn.RemoteAddr(), conn.LocalAddr(), id)
	time.Sleep(200 * time.Millisecond)

	requestNumber := 5
	conn.SetDeadline(time.Now().Add(5*time.Millisecond))

	for i:=0; i<requestNumber; i++ {
		i32Req := rand.Int31()
		n, err := write(conn, fmt.Sprintf("%d", i32Req))
		if err != nil {
			printLog("Write Error:%s (Client[%d]) \n", err, id)
			continue
		}
		printLog("Sent request (written %d bytes):%d (Client[%d]) \n", n, i32Req, id)
	}

	for j := 0;j < requestNumber; j++ {
		strResp, err := read(conn)
		if err != nil {
			if err == io.EOF {
				printLog("The connection is closed by another side. (Client[%d]) \n", id)
			} else {
				printLog("Read Error:%s (Client[%d]) \n", err, id)
			}
			break
		}
		printLog("Received response: %s (Client[%d]) \n", strResp, id)
	}
	wg.Done()
}

/**
	读操作
 */
func read(conn net.Conn) (string, error) {
	readBytes := make([]byte, 1)
	var buffer bytes.Buffer
	for {
		_,err := conn.Read(readBytes)
		if err != nil {
			return "", err
		}
		readByte := readBytes[0]
		if readByte == DELIMITER {
			break
		}
		buffer.WriteByte(readByte)
	}
	return buffer.String(), nil
}

/**
	写操作
 */
func write(conn net.Conn, content string) (int, error) {
	var buffer bytes.Buffer
	buffer.WriteString(content)
	buffer.WriteByte(DELIMITER)
	return conn.Write(buffer.Bytes())
}

/**
	写日志方法
 */
func printLog(format string, args...interface{}) {
	fmt.Printf("%d:%s", logSn, fmt.Sprintf(format, args...))
}
