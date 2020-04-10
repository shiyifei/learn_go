/**
	本示例演示painc和recover方法的用法，当两者在同一个作用域时才能起作用
	即：defer func()部分要跟抛出异常的方法在同一个作用域下recover才能起作用
 */
package practice

import (
	"errors"
	"fmt"
)

func NumDiv(a,b int) int {
	if b == 0 {
		panic(errors.New("被除数不能为0"))
		return 0
	}
	return a/b
}

func MyRecover() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("in MyRecover(), err:", err)
		}
	}()
	g(0)
	NumDiv(100, 0)
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}

