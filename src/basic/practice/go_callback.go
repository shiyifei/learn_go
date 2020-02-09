package practice

import "fmt"

type cb func(int) int

func AboutCallback() {
	testCallBack(2,callBack)
	testCallBack(3,mul)
	testCallBack(4,func(a int) int {
		fmt.Printf("在回调函数里,a=%d \n", a)
		return a*a
	})

}


func testCallBack(x int, f cb) {
	output := f(x)
	fmt.Println("回调函数执行结果为",output)
	
}

func callBack(x int) int {
	fmt.Printf("在回调函数里, x=%d \n", x)
	return x * x
}

func mul(x int) int {
	fmt.Printf("在回调函数里，x=%d\n",x)
	return 10*x
}
