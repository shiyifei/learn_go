package main

import(
	UseStruct "use_struct/practice"
)

func main() {
	emp := UseStruct.NewEmployee("wangzhongwei", 34, "wangzhongwei@163.com")
	student := UseStruct.NewStudent("wanggengke", 29, "tsinghua")

	UseStruct.PrintObj(emp)
	UseStruct.PrintObj(student)
}
