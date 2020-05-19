package practice
/**
	本示例演示如何处理struct为Nil的情况, 将struct赋值给一个接口类型的变量时，如果再判断接口变量是否为nil，要注意写法
 */

import (
	"fmt"
	"reflect"
)

 type Explodes interface {
 	Bang()
 	Boom()
 }

 type Bomb struct {}
 func(*Bomb) Bang(){}
 func(Bomb) Boom() {}

 func ProcessNil() {
 	var bomb *Bomb = nil
 	var explodes Explodes = bomb
 	fmt.Println(bomb, explodes)

 	//判断是否nil的错误写法
 	/*if explodes != nil {
 		fmt.Println("explodes is not nil!")
 		explodes.Bang()
 		explodes.Boom()	//此处会触发panic
	} else {
		fmt.Println("explodes is nil")
	}*/



 	//判断接口类型变量是否nil的正确写法 使用反射后的值是否为Nil来判断
 	if explodes != nil && !reflect.ValueOf(explodes).IsNil() {
		fmt.Println("explodes is not nil!")
		explodes.Bang()
		explodes.Boom()
	} else {
		fmt.Println("explodes is nil")
	}

 }

