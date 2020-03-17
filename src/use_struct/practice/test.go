package practice

import (
	"fmt"
	"reflect"
)

type employee struct{
	name string
	age int
	email string
}

type student struct {
	name string
	age int
	school string
}

func NewEmployee(name string, age int, email string) employee{
	emp := employee{name, age, email}
	return emp
}
func (e employee) Print() {
	fmt.Printf("employee name:%s, age:%d, email:%s \n",e.name,e.age,e.email)
}

func NewStudent(name string, age int, school string) student{
	stud := student{name,age,school}
	return stud
}
func (e student) Print() {
	fmt.Printf("employee name:%s, age:%d, email:%s \n",e.name,e.age,e.school)
}

func PrintObj(obj interface{}) {
	//objType := reflect.TypeOf(obj)
	//fmt.Println(objType)
	v := reflect.ValueOf(obj)
	/*for i:=0;i<v.NumField();i++ {
		fmt.Printf("Field type:%T, value:%v \n", v.Field(i),v.Field(i))
	}*/
	//output := v.NumMethod()
	//fmt.Println("output:",output)

	name := reflect.TypeOf(obj).Name()
	//fmt.Printf("name  type:%T, value:%s \n", name,name)
	switch name {
	case "employee":
		fmt.Printf("employee name:%s, age:%d, email:%s \n",v.FieldByName("name"),v.FieldByName("age"),v.FieldByName("email"))
	case "student":
		fmt.Printf("student name:%s,age:%d,school:%s \n",v.FieldByName("name"),v.FieldByName("age"),v.FieldByName("school"))
	}
}

