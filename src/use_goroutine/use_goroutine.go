package main

import (
	"fmt"
	goroutine "use_goroutine/practice"
)

func main() {
	goroutine.Practice1()
	fmt.Println("===end practice1=======================================================================================")
	goroutine.Practice2()
	fmt.Println("===end practice2=======================================================================================")

	goroutine.Practice3()
	fmt.Println("===end practice3=======================================================================================")
	goroutine.Practice4()
	fmt.Println("===end practice4=======================================================================================")
	goroutine.Practice5()
	fmt.Println("===end practice5=======================================================================================")
	goroutine.AtomicLoadStore()
	fmt.Println("===end atomic.load and store===========================================================================")

}
