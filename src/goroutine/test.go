/**
	如何查看goroutine的调度过程
 */
package main

import "time"

func main() {
	//启用了10个goroutine
	for i:=0;i<10;i++ {
		go func() {
			time.Sleep(5 * time.Second)
		}()
		time.Sleep(time.Second)
	}

}
