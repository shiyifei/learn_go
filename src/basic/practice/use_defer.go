package practice

import (
	"fmt"
)

func Use_defer() {
	aboutCondition1()
	aboutCondition2()
	aboutCondition3()
	ret := aboutCondition4()
	fmt.Println("call acoutCondition4(),ret:", ret)
}

/**
 * 这里的输出结果是什么呢？
 * 输出结果是0，是以defer语句所在位置处的i为依据输出值的。
 */
func aboutCondition1() {
	i := 0
	defer fmt.Println("in abountCondition1(), i=", i)
	i++
	return
}

/**
 * 这里的输出结果又是什么呢？
 * 输出结果是1，defer语句会受到所在位置的影响
 */
func aboutCondition2() {
	i := 0
	i++
	defer fmt.Println("in abountCondition2(), i=", i)
	return
}

//注意defer语句的输出顺序，是先进后出的顺序
//最终会是在循环完毕的时候才开始执行，所以顺序正好与循环进入的顺序相反。
func aboutCondition3() {
	for i := 0; i < 4; i++ {
		defer fmt.Println("in abountCondition3(), i=", i)
	}
}

//考虑一下这里的输出结果，
//程序会执行return，导致i=1
//再执行i++, i自增1，最终结果将会是2
func aboutCondition4() (i int) {
	defer func() { i++ }()
	return 1
}
