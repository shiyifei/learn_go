package main

import (
	"time"
	channel1 "use_channel/practice"          //默认类名是practice,现在取别名为 channel1
)

func main() {
	channel1.NotUseChannel()
	time.Sleep(100*time.Millisecond)  //100ms
	channel1.UseChannel()
}

