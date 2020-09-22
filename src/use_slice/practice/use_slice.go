/**
GO中的数组是值类型，赋值给另外一个变量时实际上是原始数组的一个副本，修改该变量的值不会影响原数组的值。
slice是引用类型, 是对现有数组的引用
切片持有对底层数组的引用。只要切片在内存中，数组就不能被垃圾回收。在内存管理方面，这是需要注意的。
让我们假设我们有一个非常大的数组，我们只想处理它的一小部分。
然后，我们由这个数组创建一个切片，并开始处理切片。这里需要重点注意的是，在切片引用时数组仍然存在内存中。
一种解决方法是使用 copy 函数 func copy(dst，src[]T)int 来生成一个切片的副本。
这样我们可以使用新的切片，原始数组可以被垃圾回收。
 */
package practice


import (
	"fmt"
	"reflect"
	"sort"
)

func InStepOperate() {
	//arrStr是数组类型，会自动计算长度
	arrStr := [...]string{"java", "c", "c++", "python", "c#", "basic"}
	fmt.Printf("arrStr: length=%d,cap=%d,type is: %T, value is:%v\n", len(arrStr), cap(arrStr), arrStr, arrStr)

	//多维数组的写法
	var arrMul [3][2]string
	arrMul[0][0] = "are"
	arrMul[0][1] = "you"
	arrMul[1][0] = "ok"
	arrMul[1][1] = "how"
	arrMul[2][0] = "are"
	arrMul[2][1] = "you"

	//定义一个三行两列的数组
	arr32 := [3][2]string {
		{"hello", "shiyf"},
		{"what", "are"},
		{"you", "doing"}, //注意这里的逗号不能省略
	}
	for i, v1 := range arr32 {
		for j, v2 := range v1 {
			fmt.Printf("row=%d,col=%d, %s \n", i,j, v2)
		}
	}
	fmt.Printf("\n")

	sliceStr := arrStr[1:4:5] //下标一表示起始位置，下标二表示结束位置（不包含该位置元素），下标三表示cap容量 cap=len+1,容量可以设置大于len+1
	printSlice(sliceStr)

	sliceStr[0] = "Go"
	fmt.Println("after assigning slice, arrStr:")
	sliceStr = append(sliceStr, "ruby")

	//不超容量重新赋值后，原数组内容会跟着改变
	//切片的长度是切片中的元素数，切片的容量是创建切片索引开始的底层数组的元素个数
	fmt.Printf("arrStr: length=%d,cap=%d,type is: %T, value is:%v\n", len(arrStr), cap(arrStr), arrStr, arrStr)

	printSlice(sliceStr)

	sliceStr2 := arrStr[1:4]
	fmt.Println("after assigning slice again")
	sliceStr2 = append(sliceStr2, "ruby", "php", "javascript", "pl-sql")
	fmt.Println(arrStr) //超容量重新赋值后，原数组不会跟着改变
	fmt.Println("len(sliceStr2),cap(sliceStr2)", len(sliceStr2), cap(sliceStr2))
	fmt.Println(sliceStr2)

	sliceStr3 := []string{"are", "you", "ok"}
	var sliceStr4 = make([]string, 3)
	copy(sliceStr4, sliceStr3) //copy会复制原有切片的值，修改时不影响原有元素，原有的切片可以尽快被垃圾回收掉。
	var sliceStr5 = make([]string, 5, 10)

	sliceStr4[0] = "how"
	sliceStr4[1] = "are"
	sliceStr4 = append(sliceStr4, "you")

	fmt.Println("copy sliceStr3 to sliceStr4 and update sliceStr4:")
	fmt.Print("sliceStr4:")
	printSlice(sliceStr4)
	fmt.Print("sliceStr3:")
	printSlice(sliceStr3)

	sliceStr5 = sliceStr3[0:3]
	sliceStr5 = append(sliceStr5, "do", "not")

	fmt.Print("before changing, sliceStr5:")
	printSlice(sliceStr5)

	changeSlice(sliceStr5) //切片是引用类型，在函数中改变内容的话,在函数外部能看到已经生效了。
	fmt.Print("after changing sliceStr5:")
	printSlice(sliceStr5)

	languages := []string{"C", "C++", "Java", "Go", "PHP", "Python", "Ruby"}
	already := languages[:len(languages)-2]

	//这种赋值方法可以保证尽快地将languages数组做垃圾回收操作
	var newAlready = make([]string, len(already)) //先声明同样大小的切片
	copy(newAlready, already)                     //生成一个切片的副本，原始数组就可以尽快被垃圾回收
	fmt.Println("newAlready:", newAlready)

	other := []string{"Object-C", "Swift", "Erlang"}
	newAlready = append(newAlready, other...) //追加切片的正确写法
	fmt.Println("after append other slice, newAlready:", newAlready)

	fmt.Println("before changing, other:", other)
	//使用可变参数的函数来更改切片本身的值
	change(other...)
	fmt.Println("after changing, other:", other)

	var other1 = make([]string, len(other))
	copy(other1, other)
	fmt.Printf("before change1(), pointer:%p,cap: %d,", other1, cap(other1)) //获取切片的内存地址
	fmt.Println("before change1(), other1:", other1)
	change1(other1)                                                                             //
	fmt.Printf("after change1(), pointer:%p,cap: %d,len:%d,", other1, cap(other1), len(other1)) //这里的地址仍然是以前的地址，所以修改有效，追加无效。
	fmt.Println("after change1(), other1:", other1)

	other1 = append(other1, "javascript")
	fmt.Printf("before change2(), pointer:%p,cap: %d,len:%d,", other1, cap(other1), len(other1)) //获取切片的内存地址
	fmt.Println("before change2(), other1:", other1)
	change2(&other1)                                                        //
	fmt.Printf("after change2(), pointer:%p,cap: %d,", other1, cap(other1)) //这里的地址仍然是以前的地址，所以修改有效，追加无效。
	fmt.Println("after change2(), other1:", other1)

	other1 = append(other1, "vb.net", "shell") //一旦超出原有长度，会将原有的容量翻倍
	fmt.Printf(" pointer:%p,cap: %d,len:%d \n", other1, cap(other1), len(other1))

	useAppend()

	arrInt := []int{12, 34, 1, 4, 5, 6, 7, 23, 56, 12, 34, 56}
	sort.Ints(arrInt)	//对切片排序，升序排列
	arrUnique := RemoveDuplicate(arrInt)
	fmt.Println("   arrInt:", arrInt)
	fmt.Println("arrUnique:", arrUnique)

	//将[]interface{}转换为[]int
	arrSource := make([]int, len(arrUnique))
	for i := range arrUnique {
		arrSource[i]= arrUnique[i].(int)
	}

	var target int = 66
	pos,isExist := searchSlice(arrSource, target)
	if !isExist {
		fmt.Printf("target:%d is not Exist in arrSource \n", target)
	} else {
		fmt.Printf("target:%d is Existed in arrSource, position:%d \n", target, pos)
	}

}

/*
	只修改其中的元素值会在函数体外生效
*/
func changeSlice(strSlice []string) {
	strSlice[4] = "something"
}

/*
	函数内部追加写入的数据不会在外部生效，因为append函数执行后，会导致s切片的地址发生变化，传入的切片与外部的切片已经不同
	实参实际上是对参数的值copy的副本，修改副本不会影响原来的值。
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
	s = append(s, "C#") //append函数执行后，会导致s切片的地址发生变化，传入的切片与现在的切片已经不同
	fmt.Printf("inside change1(), pointer:%p,cap: %d,", s, cap(s))
	fmt.Println("inside change1(), s:", s)
}

/**
传入变量地址，函数内部的修改将会在函数内外都生效
*/
func change2(s *[]string) {
	(*s)[0] = "Oracle"
	fmt.Printf("before append(), pointer:%p,cap: %d,len:%d, \n", *s, cap(*s), len(*s))
	*s = append(*s, "C#") //append函数执行后，会导致s切片的地址发生变化，传入的切片与现在的切片已经不同
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
	for i := 0; i < va.Len(); i++ {
		if i > 0 {
			fmt.Println("compare ", va.Index(i-1).Interface(), va.Index(i).Interface())
		}

		//判断是否与之前的值相等,如果相等则放弃
		if i > 0 && reflect.DeepEqual(va.Index(i-1).Interface(), va.Index(i).Interface()) {
			continue
		}
		ret = append(ret, va.Index(i).Interface())
	}
	return ret
}

/**
	如何查找slice中的元素
	sourceArr 是来源切片，这里假定是int类型的
	target 是查找目标
 */
 func searchSlice(sourceArr []int, target int) (int, bool)  {
 	//SearchInts在已排序的整数片中搜索x，并返回由搜索指定的索引。返回值是插入x的索引(当x不存在时可能是len(a))。切片必须按升序排序。
 	pos := sort.SearchInts(sourceArr, target)
 	fmt.Printf("sourceArr:%+v, target:%d, pos:%d \n", sourceArr, target, pos)
 	isExist := pos < len(sourceArr) && sourceArr[pos] == target
 	return pos, isExist
 }

func printSlice(arr []string) {
	fmt.Printf("length=%d,cap=%d,type is: %T, value is:%v\n", len(arr), cap(arr), arr, arr)
}
