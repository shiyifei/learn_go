package example

import "fmt"

//定义一个函数类型的变量
type add func(a, b int) int

func Test() {
	//定义一个函数变量

	a := func(input string) {
		fmt.Printf("hello, you input string:[%s] \n", input)
	}
	a("are you ok?")
	fmt.Printf("type of a is:%T \n", a)


	func(input string) {
		fmt.Printf("hello, you input string:[%s] \n", input)
	}("I am fine, thank you!")

	var methodAdd add = func(a,b int) int {
		return a+b
	}
	fmt.Println("sum of 10 and 23 is:", methodAdd(10,23))

}
