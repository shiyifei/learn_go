package main

import (
	"basic/practice"
	"fmt"
)

func CustomerRecover1() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("in CustomerRecover1(), err:", r)
		}
	}()
}

func CustomerRecover2() {
	r := recover()
	if r != nil {
		fmt.Println("in CustomerRecover2(), err:", r)
	}
}

func main() {
	practice.GetSubMenu()
	fmt.Println("================")

	practice.UseEnv()
	fmt.Println("================")
	practice.UseTime()

	fmt.Println("================")
	return
	/*practice.Go_basic()
	practice.UseType()

	fmt.Println("panic and recover=========================")
	//practice.MyRecover()
	fmt.Println("22222222=========================")

	//这种写法recover不会起作用，因为recover方法没有跟抛出异常的方法在同一个作用域内
	CustomerRecover1()

	//这种写法能起作用，defer语句中直接捕获到抛出的异常
	defer CustomerRecover2()*/

	//这种写法能捕获到NumDiv方法中抛出的异常
	/*defer func() {
		r := recover()
		if r != nil {
			fmt.Println("in defer func(), err:", r)
		}
	}()*/
	// practice.NumDiv(3, 0)
}
