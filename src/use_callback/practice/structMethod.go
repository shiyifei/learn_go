package practice


import (
	"fmt"
	"reflect"
)

type YourT1 struct {
}

func (y *YourT1) MethodBar() {
	fmt.Println("MethodBar called")
}

type YourT2 struct {
}

func (y *YourT2) MethodFoo(i int, oo string) {
	fmt.Println("MethodFoo called", i, oo)
}

func InvokeObjectMethod(object interface{}, methodName string, args ...interface{}) {
	inputs := make([]reflect.Value, len(args))
	for i,param := range args {
		inputs[i] = reflect.ValueOf(param)
	}
	reflect.ValueOf(object).MethodByName(methodName).Call(inputs)
}

func TestInvoke() {
	InvokeObjectMethod(new(YourT2), "MethodFoo", 10, "abc")
	InvokeObjectMethod(new(YourT1), "MethodBar")
}
