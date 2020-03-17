package practice
import "fmt"

type Student struct {
	name string
	gender string
	age int
}

func main() {
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
	fmt.Println("name=",std.name)
	fmt.Println("gender=", std.gender)
	fmt.Printf("age=%d\n",std.age)
}

func printInfo(std *Student){
	fmt.Printf("name=%s,gender=%s,age=%d\n", std.name,std.gender,std.age)
}
