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


/**
	defer是后进先出的，panic需要等defer技术后才会向上传递
	出现panic恐慌的时候，会先按照defer的后入先出的顺序执行，最后才会执行panic
	近期有同学遇到多次执行的时候发现panic的执行顺序不定，那么是不是因为panic与defer没有先后关系呢
	那为什么没有加recover()时候，panic执行顺序不定呢？
	defer的执行顺序肯定是FILO的，但是没有被recover的panic协程（线程）可能争夺CPU的顺序比defer快，
	所以造成了这样的情况，也可能是写缓存问题，所以对panic进行recover将其加入到defer队列中。
 */
func defer_call() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("one=",err)
		}
	}()
	defer func() {
		fmt.Println("打印前")
	}()
	defer func() {
		fmt.Println("打印中")
	}()
	defer func() {
		fmt.Println("打印后")
	}()

	panic("触发异常")
}

/**
	打印后
	打印中
	打印前
	panic:触发异常
 */
