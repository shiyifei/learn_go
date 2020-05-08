package main

import (
	"fmt"
	"use_func/example"
)

func main() {
	//可变参数可以传入多个值
	output := example.TestVariables(1, 2, 3)
	fmt.Println("output:", output)

	//可变参数不传值也是可以的
	output = example.TestVariables()
	fmt.Println("output:", output)
	fmt.Println("=========")
	example.Test()
	example.Show()
	example.FilterStudent()
}
