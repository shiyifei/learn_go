package practice

import "fmt"

func Go_Pointer() {
	var a int = 20
	var ip *int

	fmt.Printf("赋值前，ip空指针变量值为：%x\n",ip)
	ip = &a

	fmt.Printf("a 变量的存储地址是 %x \n", &a)

	fmt.Printf("ip变量的存储地址是：%x\n", ip)

	fmt.Printf("ip变量的值是:%d \n", *ip)

}
