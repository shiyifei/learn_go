package practice

import "fmt"

func Use_assign() {
	ret := testAssign1()
	fmt.Printf("desc:%s \n", ret)

	ret = testAssign2()
	fmt.Printf("desc:%s \n", ret)
}

//注意这里的desc的声明加赋值的语句
func testAssign1() string {
	desc := "prefect"
	score := 68
	if score <= 80 && score >= 60 {
		desc = "good"
		// desc := "good" 这些的写法是不允许的，因为已经声明了desc的类型
	}
	return desc
}

//建议使用先声明再使用的方式
func testAssign2() string {
	var desc string = "prefect"
	score := 68
	if score <= 80 && score >= 60 {
		desc = "good"
	}
	return desc

}
