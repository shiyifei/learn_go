package practice

import "fmt"

type Student struct {
	name   string
	gender string
	age    int
}

type People interface {
	Speak(string) string
}

type Customer struct{}

func (stu *Customer) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are not a good boy"
	} else {
		talk = "hi"
	}
	return
}

func live() People {
	var cus *Customer
	return cus
}


var firstStudent Student

func Go_struct() {
	//会打印出A还是B呢，结果是打印出B 因为cus 这种类型已经实现了接口People的方法Speak,认为cus是该接口类型的变量
	if live() == nil {
		fmt.Println("AAAAAAAAAAAA")
	} else {
		fmt.Println("BBBBBBBBBBBB")
	}


	//var peo Customer
	//peo = Customer{}
	//调用peo.Speak()时，可以是上面注释掉的赋值语句，也可以是下面的赋值语句，这两句都是合法的
	peo := &Customer{}

	//var peo People = Customer{} //该语句会导致编译不通过，提示不是同一种类型

	think := "bitch"
	fmt.Println(peo.Speak(think))
	return

	firstStudent.name = "areyouok"

	var student1 Student
	var student2 Student

	student1.name = "wangzhongwei"
	student1.gender = "male"
	student1.age = 33

	student2.name = "shiyifei"
	student2.gender = "male"
	student2.age = 35

	PrintStudent(student1)

	PrintStudent(student2)

	var ptr_std *Student

	ptr_std = &student1

	PrintStudent(*ptr_std)

	printInfo(&student2)
}

func PrintStudent(std Student) {
	fmt.Println("name=", std.name)
	fmt.Println("gender=", std.gender)
	fmt.Printf("age=%d\n", std.age)
}

func printInfo(std *Student) {
	fmt.Printf("name=%s,gender=%s,age=%d\n", std.name, std.gender, std.age)
}
