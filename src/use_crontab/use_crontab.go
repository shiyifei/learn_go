package main

import "use_crontab/practice"

func main() {
	go practice.CreateServer()
	practice.StartTimer()

	//practice.InvokeObjectMethod(new(practice.Processor), "ProcessOrder", "areyouok")
}
