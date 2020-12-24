/**
 * @Author:shiyf
 * @Date: 2020/12/25 0:21
 **/

package practice

import "fmt"

type student struct {
	Name string
	Age int
}

/**
	你会发现最终生成map实际上并不是我们想要的，里面只含有切片中的最后一个元素了
	与Java的foreach一样，for range 都是使用副本的方式
	m[stu.Name]=&stu实际上一致指向同一个指针， 最终该指针的值为遍历的最后一个struct的值拷贝
 */
func pase_student() {
	m := make(map[string]*student)
	arrStudent := []student{
		{Name:"zhao", Age:22},
		{Name:"wang", Age:25},
		{Name:"li", Age:24},
	}
	for _, stu := range arrStudent {
		m[stu.Name] = &stu
	}

	for _, stu := range m {
		fmt.Printf("stu:%+v\n", stu)
	}
}

/**
	将一个切片复制到map中的正确写法
 */
func parse_student() {
	m := make(map[string]*student)
	arrStudent := []student{
		{Name:"zhao", Age:22},
		{Name:"wang", Age:25},
		{Name:"li", Age:24},
	}

	for i:=0;i<len(arrStudent);i++ {
		m[arrStudent[i].Name] = &arrStudent[i]
	}

	for _, stu := range m {
		fmt.Printf("stu:%+v\n", stu)
	}
}