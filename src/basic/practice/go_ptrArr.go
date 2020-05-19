package practice

import "fmt"
const MAX int = 3

func Go_ptrArr() {
	a := []int{10,20,100}
	var i int
	
	var ptr [MAX]*int

	for i=0;i<MAX;i++ {
		ptr[i] = &a[i]
	}

	for i=0;i<MAX;i++ {
		fmt.Printf("ptr[%d] is:%d, addr:[%x]\n",i,*ptr[i], ptr[i])
	}

}
