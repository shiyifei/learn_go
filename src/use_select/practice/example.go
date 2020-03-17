/** select 语句用于在多个发送/接收信道中进行选择
	select 语句会一直阻塞，直到发送/接收操作准备就绪。
    如果有多个信道操作准备完毕，select会随机地选取其中之一执行
**/

package practice

import (
	"fmt"
	"time"
)

func server1(channel chan string) {
	time.Sleep(100*time.Millisecond)
	channel <- "from server1"
}

func server2(channel chan string) {
	time.Sleep(100*time.Millisecond)
	channel <- "from server2"
}

/**
	最终的输出结果是随机的
 */
func TestSelect() {
	var chan1 = make(chan string)
	var chan2 = make(chan string)
	go server1(chan1)
	go server2(chan2)

	for {
		time.Sleep(10*time.Millisecond)
		//如果有多个信道操作准备完毕，select会随机地选取其中之一执行
		select {
		case s1 := <-chan1:
			fmt.Println(s1)
			return
		case s2 := <-chan2:
			fmt.Println(s2)
			return
		default:				//设置default语句，可以避免select语句一直阻塞，但是也可能会导致没有收到正确的结果。
			fmt.Println("no value received")
		}
	}

}

