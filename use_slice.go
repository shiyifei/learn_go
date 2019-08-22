package main

import "fmt"

func main() {
	arrStr := [...]string{"java", "c","c++","python","c#","basic"}
	sliceStr := arrStr[1:4:5]  //下标一表示起始位置，下标二表示结束位置（不包含该位置元素），下标三表示cap容量 cap=len+1,容量可以设置大于len+1
	fmt.Println("len(arrStr),cap(arrStr)", len(arrStr), cap(arrStr))
	fmt.Println(arrStr)
	fmt.Println("len(sliceStr),cap(sliceStr)", len(sliceStr), cap(sliceStr))
	fmt.Println(sliceStr)

	sliceStr[0] = "Go"
	fmt.Println("after assigning slice, arrStr:")
	sliceStr = append(sliceStr, "ruby")
	fmt.Println(arrStr)  //不超容量重新赋值后之前的数组会跟着改变
	fmt.Println("len(sliceStr),cap(sliceStr)", len(sliceStr), cap(sliceStr))
	fmt.Println(sliceStr)

	sliceStr2 := arrStr[1:4]
	fmt.Println("after assigning slice again")
	sliceStr2 = append(sliceStr2, "ruby", "php", "javascript", "pl-sql")
	fmt.Println(arrStr)  //超容量重新赋值后之前的数组不会跟着改变
	fmt.Println("len(sliceStr2),cap(sliceStr2)", len(sliceStr2), cap(sliceStr2))
	fmt.Println(sliceStr2)

	sliceStr3 := []string{"are","you","ok"}
	var sliceStr4 = make([]string ,3)
	copy(sliceStr4, sliceStr3)  //copy会复制原有切片的值，修改时不影响原有元素。
	var sliceStr5 = make([]string,5,10)

	sliceStr4[0] = "how"
	sliceStr4[1] = "are"
	sliceStr4 = append(sliceStr4, "you")

	sliceStr5 = sliceStr3[0:3]
	sliceStr5 = append(sliceStr5, "do","not")

	fmt.Println(sliceStr3)
	fmt.Println(sliceStr4)

	fmt.Println("before changing, sliceStr5:")
	fmt.Println(sliceStr5)


	changeSlice(sliceStr5)  //切片是引用类型，在函数中改变内容的话,在函数外部能看到已经生效了。
	fmt.Println("after changing sliceStr5:")
	fmt.Println(sliceStr5)

	languages := []string{"C","C++","Java","Go","PHP","Python","Ruby"}
	already := languages[:len(languages)-2]

	//这种赋值方法可以保证尽快地将languages数组做垃圾回收操作
	var newAlready = make([]string, len(already))    //先声明同样大小的切片
	copy(newAlready, already)						//生成一个切片的副本，原始数组就可以尽快被垃圾回收
	fmt.Println("newAlready:",newAlready)

	other := []string{"Object-C","Swift"}
	newAlready = append(newAlready, other...)   //追加切片的正确写法
	fmt.Println("after append other slice, newAlready:", newAlready)

	//使用可变参数的函数来更改切片本身的值
	change(other...)
	fmt.Println("after changing, other:", other)

	var other1 = make([]string, len(other))
	copy(other1, other)
	change1(other1)
	fmt.Println("after 111 changing, other1:", other1)

}

func changeSlice(strSlice []string) {
	strSlice[4] = "something"
}

func change(s ...string) {
	s[0] = "Sql"
	s = append(s, "C#") //追加的元素不会在函数外部显示
	fmt.Println("inside  changing, s:", s)
}

/**
切片包含长度、容量和指向数组第零个元素的指针。
当切片传递给函数时，即使它通过值传递，指针变量也将引用相同的底层数组。
因此，当切片作为参数传递给函数时，函数内所做的更改也会在函数外可见。让我们写一个程序来检查这点。
 */
func change1(s []string) {
	s[0] = "Sql111"
	s = append(s, "C#")  //追加的元素不会在函数外部显示，考虑一下原因
	fmt.Println("inside 111 changing, s:", s)
}
