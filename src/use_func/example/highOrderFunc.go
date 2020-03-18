/**
	wiki 把高阶函数（Hiher-order Function）定义为：满足下列条件之一的函数：
		接收一个或多个函数作为参数
		返回值是一个函数
 **/

package example

import "fmt"

/*
	函数作为参数，传递给其他函数
 */
func simple(a func(a,b int) int ) {
	fmt.Println(a(112,442))
}

/*
	在其它函数中返回函数
 */
func exam() func(a ,b int) int {
	f := func(a,b int) int {
		return a+b
	}
	return f
}

/*
	闭包Closure是匿名函数的一个特例，当一个匿名函数所访问的变量的定义在函数体的外部时，就称这样的匿名函数为闭包
 */
func closure() {
	var a int = 123

	func(){
		fmt.Println("a=",a)   //这里的a变量定义在函数体的外部
	}()
}

func appendStr() func(string) string {
	t := "Hello"

	c := func(input string) string {
		t = t + " " + input
		return t
	}
	return c
}

func Show() {
	//先定义一个函数类型的变量
	funcA := func(a,b int) int {
		return a+b
	}

	//调用时，直接传入函数类型的变量
	simple(funcA)

	funcB := exam()
	result := funcB(11, 655)
	fmt.Println(result)

	closure()

	a := appendStr()   //函数类型的变量
	b := appendStr()

	fmt.Println(a("world"))
	fmt.Println(b("Everyone"))

	fmt.Println(a("Gopher"))
	fmt.Println(b("!"))
}

/** 以下是一个头等函数的实际用途 **/
type student struct {
	name string
	grade int
	score int
}

func filter(s []student, f func(student) bool ) []student {
	var list []student
	for _, v := range s {
		if f(v) == true {
			list = append(list, v)
		}
	}
	return list
}

func FilterStudent() {
	s1 := student{
		name:"wangxiaolan",
		grade:1,
		score:88,
	}
	s2 := student{
		name:"zhangxiaoming",
		grade:2,
		score:77,
	}
	s3 := student {
		name:"liuchen",
		grade:3,
		score:87,
	}
	s := []student{s1,s2,s3}

	inFunc := func(source student) bool {
		return source.grade ==3 && source.score>80
	}
	result := filter(s, inFunc)
	fmt.Println(result)
}

