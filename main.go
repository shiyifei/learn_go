package main

import (
	"fmt"
	"hello"
	"unsafe"
	"project1"
)

var age int = 30
var username string = "shiyf"
var score float32 = 1189.5
var isOk bool = false
var a, b, c int = 4, 5, 6

type Person struct {
	name string
	age  int
}

//var P = Person{"huqilong", 36}
var P = Person{age:32}

const LENGTH int = 5
const WIDTH = 10

const (
	Unknown = 0
	Male    = 1
	Famale  = 2
)

type T struct {
	t1 byte
	t2 int32
	t3 int64
	t4 string
	t5 bool
}

func main() {
	fmt.Printf("Hello,%s,your score is:%.2f,what are you doing now?age:%d,are you ok?%t \n",
		username, score, age, isOk)

	//unsafe.Sizeof(age)返回占用字节数，32位系统下大小是 32 位（4 字节）, 64位系统下，age会占用 64 位（8 字节）的大小
	fmt.Printf("type of age is:%T, size of age is:%d \n", age, unsafe.Sizeof(age))

	a, b := 6, 8 //a,b,c 局部赋值优先

	// a,b,c := 1,2,3
	fmt.Println(a, b, c)

	fmt.Printf("person name:%s,age:%d \n", P.name, P.age)

	//P := Person{name:"wanggengke"}
	var P Person
	P.name = "areyouok"
	P.age = 33
	fmt.Printf("person name:%s,age:%d \n", P.name, P.age)
	fmt.Println(P)

	var _, ret, retStr = numbers()
	fmt.Println(ret, retStr)

	var area int
	area = LENGTH * WIDTH
	fmt.Printf("area is:%d \n", area)

	fmt.Printf("male is:%d \n", Male)

	var str string = "abc"
	fmt.Println(str, len(str), unsafe.Sizeof(str))

	//unsafe包其实是指针
	fmt.Println("----------unsafe.Pointer---------")
	t := &T {1,2,3, "this is an example", true}
	ptr := unsafe.Pointer(t)  //获取变量t的通用指针
	t1 := (*byte)(ptr)		//unsafe.Pointer可以和普通指针进行相互转换
	fmt.Println(*t1)	 	//t.t1当前的值

	//unsafe.Pointer 可以和 uintptr 进行相互转换 uintptr(ptr)
	t2 := (*int32)(unsafe.Pointer(uintptr(ptr) + unsafe.Offsetof(t.t2)))  	//指针移动到t.t2位置
	fmt.Println(*t2)  //t.t2当前的值
	*t2 = 99	//实际上会更改t.t2的值
	fmt.Println(t) //可以直接打印一个stuct对象,其中的元素会用空格分开

	t3 := (*int64)(unsafe.Pointer(uintptr(ptr) + unsafe.Offsetof(t.t3)))	//指针移动到t.t3位置
	fmt.Println(*t3)  //t.t3当前的值
	*t3 = 123  //实际上会更改t.t3的值
	fmt.Println(t)


	numC := project1.IntAdd(a, b)
	fmt.Println("numC=", numC)

	numD := hello.GetArea(a, b)
	fmt.Println("numD=", numD)
}

func numbers() (int, int, string) {
	a, b, c := 1, 2, "are you ok?"
	return a, b, c
}
