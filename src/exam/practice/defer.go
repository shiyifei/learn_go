package practice

import "fmt"

func f1() (t int) {
	t = 5
	defer func() {
		t = t + 5
	}()
	return t
}

func f2() (r int) {
	defer func(t int) int {
		t = t + 5
		return t
	}(r)
	return 1
}

func f2Other() (r int) {
	r = 1
	func(t int) int{
		t = t +5
		return t
	}(r)
	return
}

func f3(j int) int {
	j = 1

	defer func() {
		j++
		fmt.Println("222 j=",&j, j)
	}()
	fmt.Println("111 j=",&j, j)
	return j
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
	c := f3(0)
	fmt.Println("after f3(), c=", &c, c)

	aboutCondition1()
	aboutCondition2()
	aboutCondition3()
}


