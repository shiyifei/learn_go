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
	fmt.Println(sliceStr5)

	fmt.Println("after changing sliceStr5:")
	changeSlice(sliceStr5)  //切片是引用类型，在函数中改变内容的话,在函数外部能看到已经生效了。
	fmt.Println(sliceStr5)

}

func changeSlice(strSlice []string) {
	strSlice[4] = "something"
}
