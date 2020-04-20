/**
map是引用类型，重新赋值后修改其中元素的值会影响之前的变量的值。
*/
package practice

import "fmt"

func UseMap() {
	var empSalary map[string]int

	empSalary = make(map[string]int)

	empSalary["john"] = 12000
	empSalary["herry"] = 15000
	empSalary["steve"] = 18000

	fmt.Println("content of empSalary:", empSalary) //map可以直接打印出来，形如：[key1:value1 key2:value2]

	var emp string = "john"
	fmt.Printf("salary of %s is:%d\n", emp, empSalary[emp])

	emp = "tom"
	fmt.Printf("salary of %s is:%d\n", emp, empSalary[emp]) //直接取值会获取到默认值，int类型默认值为0，如果正好有值为0的，就区分不出到底本来值就是0还是根本不存在该元素
	salary, ok := empSalary[emp]                            //判断元素是否存在，然后再取值，这样取出的值才是正确的。
	if ok == true {
		fmt.Printf("salary of %s is:%d\n", emp, salary)
	} else {
		fmt.Println(emp, " not found")
	}
	//删除map中的元素
	delete(empSalary, "john")

	changeMap(empSalary)

	// 用 for range循环遍历map中的所有元素,当使用 for range 遍历 map 时，不保证每次执行程序获取的元素顺序相同。
	for key, value := range empSalary {
		fmt.Printf("empSalary[%s]:%d \n", key, value)
	}

	newSalary := map[string]int{
		"john":  12000,
		"steve": 15000,
		"herry": 18000,
	}

	isEqual := compareMap(empSalary, newSalary)
	if isEqual == true {
		fmt.Println("empSalary is Equal newSalary")
	} else {
		fmt.Println("empSalary is not equal newSalary")
	}

}

/**
由于map是引用类型，函数中的参数如果传入map类型值，在函数内修改元素值的话，在函数外也会生效。
*/
func changeMap(input map[string]int) {
	input["steve"] = 19000
	input["brown"] = 22000
}

/**
比较两个map变量的值是否相等的方法
*/
func compareMap(map1, map2 map[string]int) bool {
	for key, value := range map1 {
		if map2[key] != value {
			return false
		}
	}
	return true
}
