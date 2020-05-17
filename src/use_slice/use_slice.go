package main

import (
	"fmt"
	UseSlice "use_slice/practice"
)

func main() {
	UseSlice.TestSlice()
	fmt.Println("======end test slice()==============")

	UseSlice.BasicOperate()
	fmt.Println("====end basic operation======================================================================================")
	UseSlice.InStepOperate()
	fmt.Println("====end in step operation======================================================================================")
}
