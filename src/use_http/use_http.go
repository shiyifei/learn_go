package main

import (
	"fmt"
	"use_http/practice"
)

func main() {
	practice.SendRequest()
	fmt.Println("1111111==================")
	practice.HttpDo()
	fmt.Println("222222==================")
	//practice.HttpPost()
	//fmt.Println("333333==================")
	//practice.HttpPostForm()
	//fmt.Println("444444==================")
	//practice.CreateServer()

}
