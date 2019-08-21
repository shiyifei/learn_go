package main

import (
	"fmt"
	"hello"
)

func main() {
	fmt.Println("Hello, shiyifei, what are you doing now ?")

	var a,b int
	a = 20
	b = 30

	area := hello.GetArea(a, b)

	fmt.Println("area is:", area)

}
