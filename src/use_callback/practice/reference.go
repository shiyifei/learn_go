package practice

import (
	"fmt"
	"reflect"
)

func Test(){
	ret1 := Apply(Hello, []interface{}{2, 3})
	for _, v:= range ret1{
		fmt.Println(v)
	}
	ret2 := Apply(Format, []interface{}{"hello", 44})
	for _, v:= range ret2{
		fmt.Println(v)
	}
	ret3 := Apply("processOrder", []interface{}{"shiyf"})
	for _, v:= range ret3{
		fmt.Println(v)
	}
}

func Apply(f interface{}, args []interface{})[]reflect.Value{
	fun := reflect.ValueOf(f)
	in := make([]reflect.Value, len(args))
	for k,param := range args{
		in[k] = reflect.ValueOf(param)
	}
	r := fun.Call(in)
	return r
}

// 变参
func Format(a ...interface{}) string{
	fmt.Println(a)
	return "format return"
}


func Hello(a int, b int)(int, string){
	return a + b, "hello"
}

func processOrder(input string) string {
	return "hello,"+input
}
