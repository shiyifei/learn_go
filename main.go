package main

import (
	"fmt"
	"hello/other"
	"unsafe"
	"project1"
)

var age int32 = 30
var username string = "shiyf"
var score float32 = 1189.5
var isOk bool = false
var a, b, c int = 4, 5, 6

type Person struct {
	name string
	age  int
}

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
	fmt.Printf("Hello,%s,your score s:%.2f,what are you doing now?age:%d,are you ok?%t \n",
		username, score, age, isOk)

	score = score + float32(age)  //强制类型转换，不然会报错
	//unsafe.Sizeof(score)返回占用字节数
	fmt.Printf("score is:%.2f, type of score is:%T, size of score is:%d \n", score, score, unsafe.Sizeof(score))


	//unsafe.Sizeof(age)返回占用字节数，32位系统下大小是 32 位（4 字节）, 64位系统下，age会占用 64 位（8 字节）的大小
	fmt.Printf("type of age is:%T, size of age is:%d \n", age, unsafe.Sizeof(age))

	const strName string = "shiyf,what are you doing now?"
	fmt.Printf("value of strName:%s,type of strName is:%T, size of strName is:%d \n", strName, strName, unsafe.Sizeof(strName))

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

	var _, ret, retStr = numbers()  //调用多个返回值的方法
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

	numD := other.GetArea(a, b)
	fmt.Println("numD=", numD)

	var len float64 = 20.0
	var width float64 = 30.0
	var f11, f12 float64
	f11,f12 = rectProps(len, width)  //调用多个返回值的方法

	fmt.Println("area=",f11,",perimeter=",f12)

	printRows()

	useSwitch()
}

/**
返回多个值的建议写法,比较好理解
 */
func numbers() (int, int, string) {
	a, b, c := 1, 2, "are you ok?"
	return a, b, c
}

/**
返回多个值的方法，不建议这样写，因为不好理解
 */
func rectProps(length,width float64)(area, perimeter float64) {
	area = length * width
	perimeter = 2*(length+width)
	return
}


func printRows() {
	for i:=10;i<20;i++ {
		for j:=1;j<=10;j++ {
			if  j-1 == i-10 {
				fmt.Printf("%d * %d = %d \n", i, j, i*j)
			}
		}
	}
}

func useSwitch() {
	letter := "i"
	//switch不需要在每个case语句后写break;
	switch letter {
	case "a","e","i","o","u":
		fmt.Println("这是一个元音字母")
	default:
		fmt.Println("这个一个非元音字母")
	}
}