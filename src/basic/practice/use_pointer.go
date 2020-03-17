package practice

import (
	"fmt"
	"unsafe"
)

func main() {
	var f32 float32 = 3.1415926
	pointer := unsafe.Pointer(&f32)
	fmt.Println(pointer)

	vptr := (*float32)(pointer)
	fmt.Println(vptr)
	fmt.Println(*vptr)

	output := (*int32)(pointer)
	fmt.Println(output)
	fmt.Println(*output)

}
