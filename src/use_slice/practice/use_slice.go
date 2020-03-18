package practice

import (
	"fmt"
	"reflect"
	"sort"
)

func InStepOperate() {
	arrStr := [...]string{"java", "c","c++","python","c#","basic"}
	fmt.Printf("arrStr: length=%d,cap=%d,type is: %T, value is:%v\n", len(arrStr), cap(arrStr), arrStr, arrStr)
	sliceStr := arrStr[1:4:5]  //下标一表示起始位置，下标二表示结束位置（不包含该位置元素），下标三表示cap容量 cap=len+1,容量可以设置大于len+1
	printSlice(sliceStr)

	sliceStr[0] = "Go"
	fmt.Println("after assigning slice, arrStr:")
	sliceStr = append(sliceStr, "ruby")

	//不超容量重新赋值后之前的数组会跟着改变
	fmt.Printf("arrStr: length=%d,cap=%d,type is: %T, value is:%v\n", len(arrStr), cap(arrStr), arrStr, arrStr)

	printSlice(sliceStr)

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

	fmt.Println("copy sliceStr3 to sliceStr4 and update sliceStr4:")
	fmt.Print("sliceStr4:")
	printSlice(sliceStr4)
	fmt.Print("sliceStr3:")
	printSlice(sliceStr3)

	sliceStr5 = sliceStr3[0:3]
	sliceStr5 = append(sliceStr5, "do","not")

	fmt.Print("before changing, sliceStr5:")
	printSlice(sliceStr5)


	changeSlice(sliceStr5)  //切片是引用类型，在函数中改变内容的话,在函数外部能看到已经生效了。
	fmt.Print("after changing sliceStr5:")
	printSlice(sliceStr5)


	languages := []string{"C","C++","Java","Go","PHP","Python","Ruby"}
	already := languages[:len(languages)-2]

	//这种赋值方法可以保证尽快地将languages数组做垃圾回收操作
	var newAlready = make([]string, len(already))    //先声明同样大小的切片
	copy(newAlready, already)						//生成一个切片的副本，原始数组就可以尽快被垃圾回收
	fmt.Println("newAlready:",newAlready)

	other := []string{"Object-C","Swift", "Erlang"}
	newAlready = append(newAlready, other...)   //追加切片的正确写法
	fmt.Println("after append other slice, newAlready:", newAlready)

	fmt.Println("before changing, other:", other)
	//使用可变参数的函数来更改切片本身的值
	change(other...)
	fmt.Println("after changing, other:", other)

	var other1 = make([]string, len(other))
	copy(other1, other)
	fmt.Printf("before change1(), pointer:%p,cap: %d,", other1, cap(other1)) //获取切片的内存地址
	fmt.Println("before change1(), other1:", other1)
	change1(other1)  //
	fmt.Printf("after change1(), pointer:%p,cap: %d,len:%d,", other1, cap(other1), len(other1)) //这里的地址仍然是以前的地址，所以修改有效，追加无效。
	fmt.Println("after change1(), other1:", other1)


	other1 = append(other1, "javascript")
	fmt.Printf("before change2(), pointer:%p,cap: %d,len:%d,", other1, cap(other1), len(other1)) //获取切片的内存地址
	fmt.Println("before change2(), other1:", other1)
	change2(&other1)  //
	fmt.Printf("after change2(), pointer:%p,cap: %d,", other1, cap(other1)) //这里的地址仍然是以前的地址，所以修改有效，追加无效。
	fmt.Println("after change2(), other1:", other1)

	other1 = append(other1, "vb.net", "shell") //一旦超出原有长度，会将原有的容量翻倍
	fmt.Printf(" pointer:%p,cap: %d,len:%d \n", other1, cap(other1), len(other1))

	useAppend()

	arrInt := []int{12,34,1,4,5,6,7,23,56,12,34,56}
	sort.Ints(arrInt)
	arrUnique := RemoveDuplicate(arrInt)
	fmt.Println("   arrInt:", arrInt)
	fmt.Println("arrUnique:", arrUnique)

}

/*
	只修改其中的元素值会在函数体外生效
 */
func changeSlice(strSlice []string) {
	strSlice[4] = "something"
}

/*
	函数内部追加写入的数据不会在外部生效，因为append函数执行后，会导致s切片的地址发生变化，传入的切片与外部的切片已经不同
 */
func change(s ...string) {
	s[0] = "Basic"
	s = append(s, "C#") //追加的元素不会在函数外部显示
	fmt.Println("inside change(), s:", s)
}

/*
	切片包含长度、容量和指向数组第零个元素的指针。
	当切片传递给函数时，即使它通过值传递，指针变量也将引用相同的底层数组。
	因此，当切片作为参数传递给函数时，函数内所做的更改也会在函数外可见。让我们写一个程序来检查这点。
 */
func change1(s []string) {
	s[0] = "Sql"
	fmt.Printf("before append(), pointer:%p,cap: %d \n", s, cap(s))
	s = append(s, "C#")  //append函数执行后，会导致s切片的地址发生变化，传入的切片与现在的切片已经不同
	fmt.Printf("inside change1(), pointer:%p,cap: %d,", s, cap(s))
	fmt.Println("inside change1(), s:", s)
}

/**
	传入变量地址，函数内部的修改将会在函数内外都生效
 */
func change2(s *[]string) {
	(*s)[0] = "Oracle"
	fmt.Printf("before append(), pointer:%p,cap: %d,len:%d, \n", *s, cap(*s), len(*s))
	*s = append(*s, "C#")  //append函数执行后，会导致s切片的地址发生变化，传入的切片与现在的切片已经不同
	fmt.Printf("inside change2(), pointer:%p,cap: %d,len:%d,", *s, cap(*s), len(*s))
	fmt.Println("inside change2(), s:", *s)
}

/*
	向切片中增加元素时会导致容量(cap)成倍递增
 */
func useAppend() {
	var numbers []int
	for i := 0; i < 10; i++ {
		numbers = append(numbers, i)
		fmt.Printf("len: %d  cap: %d pointer: %p\n", len(numbers), cap(numbers), numbers)
	}
}

/*
	slice中元素去重
 */
func RemoveDuplicate(s interface{}) (ret []interface{}) {
	if reflect.TypeOf(s).Kind() != reflect.Slice {
		return ret
	}

	va := reflect.ValueOf(s)
	for i:=0;i<va.Len();i++ {
		if i>0 {
			fmt.Println("compare ",va.Index(i-1).Interface(), va.Index(i).Interface())
		}

		//判断是否与之前的值相等,如果相等则放弃
		if i>0 && reflect.DeepEqual(va.Index(i-1).Interface(), va.Index(i).Interface()) {
			continue
		}
		ret = append(ret, va.Index(i).Interface())
	}
	return ret
}



func printSlice(arr []string) {
	fmt.Printf("length=%d,cap=%d,type is: %T, value is:%v\n", len(arr), cap(arr), arr, arr)
}