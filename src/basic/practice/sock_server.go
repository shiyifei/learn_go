package practice

import (
	"bytes"
	"fmt"
	"io"
	"math"
	"net"
	"strconv"
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
	go doWork(wg)
	time.Sleep(500* time.Millisecond)
	wg.Wait()
}

/**
	初始化socket、监听并处理读写事件
 */
func doWork(wg sync.WaitGroup) {
	var listener net.Listener
	listener, err := net.Listen(SERVER_NETWORK, SERVER_ADDRESS)
	if err != nil {
		printLog("Listen Error:%s \n", err)
	}
	defer listener.Close()
	printLog("Got listener for the server.(local address:%s)\n", listener.Addr())
	for {
		conn, err:= listener.Accept()
		if err != nil {
			printLog("Accept Error:%s \n", err)
		}
		printLog("Established a connection with a client application. (remote address:%s) \n", conn.RemoteAddr())
		go handleConn(conn)
	}
	wg.Done()
}

/**
	针对Socket做读写操作
 */
func handleConn(conn net.Conn) {
	for {
		conn.SetReadDeadline(time.Now().Add(10 * time.Second))
		//read message from socket
		strReq, err := read(conn)
		if err != nil {
			if err == io.EOF {
				printLog("The connection is closed by another side.(Server) \n")
			} else {
				printLog("Read Error: %s (Server) \n", err)
			}
			break
		}
		printLog("Received request :%s (Server) \n", strReq)

		//processing received message sent by client
		i32Req, err := strconv.Atoi(strReq)
		if err != nil {
			n, err := write(conn, err.Error())
			if err != nil {
				printLog("Write Error (written %d bytes):%s (Server) \n", err)
			}
			printLog("Sent response (written %d bytes):%s (Server) \n", n, err)
		}

		//send message to client
		f64Resp := math.Cbrt(float64(i32Req))
		respMsg := fmt.Sprintf("The cube root of %d is %f.", i32Req, f64Resp)
		n, err := write(conn, respMsg)
		if err != nil {
			printLog("Write Error:%s (Server) \n", err)
		}
		printLog("Sent response (writted %d bytes): %s (Server)\n", n, respMsg)
	}
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