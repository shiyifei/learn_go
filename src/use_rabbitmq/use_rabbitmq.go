package main

import "use_rabbitmq/practice"

func main() {
	//practice.Produce()
	//
	//practice.Consume()

	practice.GenerateTask([]string{"aaa","bbb","ccc", "ddd"})

	practice.ProcessTask()
	practice.ProcessTask()

}
