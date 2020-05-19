package main

import (
	"fmt"
	"reflect"
	"use_interface/practice"
)

func returnInterface() interface{} {
	return "are you ok, shiyifei?"
}

type Student struct {
	name string
	class string
}

func main() {
	practice.ProcessNil()
	fmt.Println("===========end process nil=======================")
	return


	practice.TestConvert()

	var varA interface{}
	varA = "are you ok?"
	fmt.Println(reflect.TypeOf(varA), "==========111=======")
	fmt.Println(varA,  "==========111=======")

	result, ret := varA.(string)
	fmt.Println(reflect.TypeOf(result), "==========112=======")
	fmt.Println(result, ret, "==========112=======")

	//var valB practice.MyString
	valRet := returnInterface()
	fmt.Println(reflect.TypeOf(valRet), "==========221=======")
	fmt.Println(valRet, ret, "==========221=======")

	valB, ret := valRet.(string)
	fmt.Println(reflect.TypeOf(valB), "==========222=======")
	fmt.Println(valB, ret, "==========222=======")

	valRet = returnInterface()
	valC, ret := valRet.(practice.MyString)
	fmt.Println(reflect.TypeOf(valC), "==========333=======")
	fmt.Println(valC, ret, "==========333=======")

	valRet = returnInterface()
	valD, ret := valRet.(*Student)
	fmt.Println(reflect.TypeOf(valD), "==========444=======")
	fmt.Println(valD, ret, "==========444=======")

	//实例1
	var input practice.MyString
	input = practice.MyString("shiyf, what are you doing now?")
	var v practice.VowersFinder
	v = input  //直接可以将input赋值给接口变量

	fmt.Printf("Vowers are %c\n", v.FindVowers()) //直接调用接口变量的方法,打印的是用空格分开的字符数组
	fmt.Println(practice.MyString(v.FindVowers()))					//打印的是字符串


	//实例2， 模拟针对不同类型的员工薪资计算公司的总费用支出
	emp1 := practice.Permanent{1, 5000, 20}
	emp2 := practice.Permanent{2, 6000, 30}
	emp3 := practice.Contract{3, 3000}

	employees := []practice.SalaryCalculator{emp1, emp2, emp3}  //这里构造一个接口类型的切片变量并初始化
	total := practice.TotalExpense(employees)
	fmt.Printf("Total expense Per Month is $%d\n", total)


	//实例3，模拟自定义类型实现接口，以指针的方式
	var sortableStrs practice.SortableStrs
	sortableStrs = practice.SortableStrs{"b1", "a2", "c4"}
	_,ok := interface{}(&sortableStrs).(practice.Sortable)
	fmt.Println(ok)

	_,ok = interface{}(&sortableStrs).(practice.SortInterface)
	fmt.Println(ok)

	sortableStrs.Sort()
	fmt.Printf("after sort(), %v \n", sortableStrs)

}
