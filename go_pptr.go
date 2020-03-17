package main
import "fmt"

func main(){
	var a int
	var ptr *int
	var pptr **int

	a=300

	ptr = &a

	pptr = &ptr

	fmt.Println("a=",a)
	fmt.Println("ptr的变量值为", *ptr)
	fmt.Println("pptr的变量值为",**pptr)
}

