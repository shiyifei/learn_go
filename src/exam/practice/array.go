package practice

import (
	"fmt"
)

/**
	在数组里获取两数之和等于目标值的两个数
 */
func GetNumber(arrInt [5]int, target int) (int, int) {
	arrLen := len(arrInt)
	if arrLen == 0 {
		return -1, -1
	}
	for i := 0; i < arrLen; i++ {
		for j := i + 1; j < arrLen; j++ {
			if arrInt[i]+arrInt[j] == target {
				return arrInt[i], arrInt[j]
			}
		}
	}
	return -1, -1
}

/**
	打印九九乘法表
 */
func PrintMulti() {
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			/*num := i * j
			ret := strconv.Itoa(num)
			if num < 10 {
				ret = strconv.Itoa(num) + " "
			}*/
			fmt.Printf("%d*%d=%2d ", i, j, i*j)
		}
		fmt.Println()
	}
}

/**
	思考题
 */
func Test() {
	var numbers = [...] int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	myslice := numbers[4:6]
	fmt.Printf("myslice为 %d, 其长度是%d, 其容量是%d \n", myslice, len(myslice), cap(myslice))

	//由于slice引用了源数组的地址，才会访问到源数组的元素
	myslice = myslice[:cap(myslice)]
	fmt.Printf("myslice的第四个元素为：%d \n", myslice[3])

	//修改第四个元素后，看源数组有无变化
	myslice[3] = 88
	//结果是会有变化的，说明这两个变量共用一个内存地址，就像指针一样。
	fmt.Println("numbers is:", numbers)

	//再追加一个元素后，如果再修改第四个元素，会发现原数组值是不变的。 追加元素后myslice占用一个新地址，与原来互不影响。
	myslice = append(myslice, 12,34,56)
	fmt.Println("myslice is:",myslice)
	myslice[3] = 66
	fmt.Println("numbers is", numbers)
}
