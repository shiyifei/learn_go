/**
	测试interface类型的强制装换，转换成功的话则会变成对应的类型
 */
package practice

import (
	"fmt"
	"reflect"
)

type user struct {
	name string
	age int8
}


func TestConvert() {
	printUser(&user{name:"wangzhongwei", age:35})
	printUser("areyouok")
}

func printUser(p interface{}) {
	user, ok := p.(*user)
	fmt.Println(reflect.TypeOf(user))
	if ok {
		fmt.Printf("name:%s, age:%d \n", user.name, user.age)
	} else {
		fmt.Println(user, "is not User type")
	}
}
