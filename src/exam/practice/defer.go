package practice

import "fmt"

func f1() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func f2() (r int) {
	defer func(r int) int {
		r = r + 5
		return r
	}(r)
	return 1
}

func f2Other() (r int) {
	r = 1
	func(r int) {
		r = r +5
	}(r)
	return
}

func f3() (i int) {
	defer func() { i++ }()
	return 1
}

func aboutCondition1() {
	i := 0
	defer fmt.Println("in abountCondition1(), i=", i)
	i++
	return
}

func aboutCondition2() {
	i := 0
	i++
	defer fmt.Println("in abountCondition2(), i=", i)
	return
}

func aboutCondition3() {
	for i := 0; i < 4; i++ {
		defer fmt.Println("in abountCondition3(), i=", i)
	}
}


func TestDefer() {
	a := f1()
	fmt.Println("after f1(), a=", a)
	b := f2()
	fmt.Println("after f2(), b=", b)
	b1 := f2Other()
	fmt.Println("after f2Other(), b1=", b1)
	c := f3()
	fmt.Println("after f3(), c=", c)

	aboutCondition1()
	aboutCondition2()
	aboutCondition3()
}


