package main

import (
	"fmt"
	"hello/other"
	"reflect"
)

func main() {
	fmt.Println("Hello, shiyifei, what are you doing now ?")

	var aa,ab int
	aa = 20
	ab = 30

	area := other.GetArea(a, b)
	fmt.Println("area is:", area)


	var a,b,c,d bool

	a = true
	b = false
	c = false
	d = false

	var output1 = a || b || !c && d
	var output2 = (( a || b ) || !c) && d
	var output3 = a || !c && d
	var output4 = b || !c && d

	var output5 = true || false && true
	var output6 = !(true || false && true)
	var output7 = true || (false && true)

	fmt.Println(output1, output2, output3, output4, output5, output6, output7)


	var numA int8

	numA = 15

	var output, Ok = interface{}(numA).(int)

	var outputA = reflect.TypeOf(numA)
	fmt.Println(output, Ok, outputA)

	var mapA map[string]int
	mapA = make(map[string]int)
	mapA["a"] = 1
	mapA["b"] = 2

	var ret,ok = mapA["a"]
	fmt.Println(ret, ok)

	ret, ok = mapA["c"]
	fmt.Println(ret, ok)

	numB := 10
	f := func() int {numB *=2; return 5}
	sliceA := []int{numB, f()}
	fmt.Println(sliceA)
}
