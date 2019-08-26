package main

import "fmt"

type Employee struct {
	Name string
	Age int
}


func main() {
	var emp Employee
	emp.Name = "wanggengke"
	emp.Age = 35

	fmt.Println("before changeName(), emp.Name=",emp.Name)
	emp.changeName("huqilong")  //值接收器的方法不能改变结果变量的值
	fmt.Println("after changeName(), emp.Name=",emp.Name)

	fmt.Println("before changeAge(), emp.Age=",emp.Age)
	(&emp).changeAge(36)
	fmt.Println("after changeAge(), emp.Age=",emp.Age)

	fmt.Println("before changeAge(), emp.Age=",emp.Age)
	emp.changeAge(37)  //系统会自动将emp转为指针
	fmt.Println("after changeAge(), emp.Age=",emp.Age)
}
/**
	使用值接收器的方法
 */
func (emp Employee)changeName(newName string) {
	emp.Name = newName
}

/**
	使用指针接收器的方法
 */
func (emp *Employee)changeAge(age int) {
	emp.Age = age
}

