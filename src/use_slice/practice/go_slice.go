package practice

import (
	"fmt"
	"sort"
	"strings"
)

type Book struct {
	name    string
	author  string
	pubtime string
	price   float32
}

type Role struct {
	Roleid   int
	Parentid int
}

//测试切片追加值是否影响原切片的值
func SliceExam1() {
	slice := make([]int, 5, 5)
	slice[0] =1
	slice[1] =2
	changeIt(slice...) //追加的元素值实际上已经超出了原切片的最大容量
	fmt.Println("after changeIt(), slice:",slice)

	changeIt(slice[0:2]...)
	fmt.Println("after changeIt slice[0:2]:",slice,",len:",len(slice),",cap:",cap(slice))
}

func changeIt(s ...int) {
	s = append(s, 3)  //已经超出原切片s的容量，但是并没有返回，并没有改变原切片的值
	fmt.Printf("in changeIt(), s:%+v, len:%d, cap:%d \n",s, len(s), cap(s))
}

//测试for range数组的机制
func SliceExam2() {
	var a = [5]int{1,2,3,4,5}
	var b [5]int

	//其实for range只是针对数组的副本进行的操作，遍历或修改v值， 并不影响a
	for i,v := range a {
		if i == 0 {
			a[1] = 12	//这种赋值语句是生效的
			a[2] = 13
 		}
		fmt.Printf("in for range(){}, i=%d, a:%+v \n", i, a)
		b[i] = v
		v = v*2
	}

	fmt.Println("a:",a)		//{1,12,13,4,5}
	fmt.Println("b:", b)	//{1,2,3,4,5}

}


//有一个上下级关系的数组
//要求根据某一个roleid找到其所有子节点，要求多级子节点都要查询出来
func TestSlice() {
	var roles map[int]Role
	roles = make(map[int]Role)
	role := Role{Roleid: 270, Parentid: 16}
	roles[270] = role
	role = Role{Roleid: 274, Parentid: 17}
	roles[274] = role
	role = Role{Roleid: 282, Parentid: 270}
	roles[282] = role
	role = Role{Roleid: 283, Parentid: 270}
	roles[283] = role
	role = Role{Roleid: 285, Parentid: 270}
	roles[285] = role
	role = Role{Roleid: 286, Parentid: 270}
	roles[286] = role
	role = Role{Roleid: 292, Parentid: 282}
	roles[292] = role
	role = Role{Roleid: 293, Parentid: 285}
	roles[293] = role

	fmt.Println("roles:", roles)

	//针对map的key进行排序,如果不排序的话可能最终结果输出顺序会出错
	var keys []int
	for k := range roles {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	currSystemRoles := make([]int, 0)
	currSystemRoles = append(currSystemRoles, 270) //先追加一个初始节点

	//如果直接遍历roles,不能保证遍历的顺序是按照由小到大输出的，但是使用排序后的keys是可以有序遍历的。
	for _, k := range keys {
		fmt.Printf("v:%+v \n", roles[k])
		for _, v1 := range currSystemRoles {
			if roles[k].Parentid == v1 {
				fmt.Printf("k=%d, v.Parentid=%d,v1=%d \n", k, roles[k].Parentid, v1)
				currSystemRoles = append(currSystemRoles, k)
				break
			}
		}
	}

	fmt.Println("currSystemRoles:", currSystemRoles)

	//这里相当于 php的 implode(",", arrayInt)
	joinStr := strings.Replace(strings.Trim(fmt.Sprint(currSystemRoles), "[]"), " ", ",", -1)
	fmt.Println("joinStr:", joinStr)
}

func BasicOperate() {
	//未设置长度的数组即为切片(slice)
	numbers := []int{1, 2, 4, 5, 6, 7}
	printArr(numbers)

	//切片赋值
	var arrInt = make([]int, 3, 10)
	arrInt[0] = 11
	arrInt[1] = 12
	arrInt[2] = 13
	printArr(arrInt)

	if arrInt == nil {
		fmt.Println("arrInt is Empty")
	}
	//从数组中截取一段作为切片的值
	arrInt = numbers[1:4] //截取数组， 会包含下标为1,2,3的数据 4-1=3，长度应该为3
	printArr(arrInt)

	arrObj := numbers[0:5]
	printArr(arrObj)

	//如何删除切片中的单个元素
	var indexNeedDel = 2
	arrObj = append(arrObj[:indexNeedDel], arrObj[indexNeedDel+1:]...)
	fmt.Printf("after deleting index 2:")
	printArr(arrObj)

	var arrStr [3]string
	arrStr[0] = "are you ok ?"
	arrStr[1] = "hello"
	arrStr[2] = "how do you do"
	fmt.Println(arrStr) //数组可以直接输出，中间用空格分开，如果元素中有空格就会导致区分不出每一个元素。

	var book1 Book
	book1.name = "php"
	book1.author = "shiyf"
	book1.pubtime = "2019-04-21"
	book1.price = 57.2

	var book2 Book
	book2.name = "java"
	book2.author = "maliqun"
	book2.pubtime = "2018-01-01"
	book2.price = 45.80

	var book3 Book
	book3.name = "go"
	book3.author = "wangjn"
	book3.pubtime = "2018-02-02"
	book3.price = 55.5

	//对象数组的赋值
	var arrBook []Book
	arrBook = append(arrBook, book1)
	arrBook = append(arrBook, book2)
	arrBook = append(arrBook, book3)
	printBookArr(arrBook)

	fmt.Printf("type of arrBook:%T \n", arrBook)

	bookArr := arrBook
	bookArr[0].author = "shi_yi_fei"
	fmt.Println("after assign and modifing books[0], arrBook:")
	printBookArr(arrBook)

	//对象切片
	var books = make([]Book, len(arrBook), 2*cap(arrBook))
	copy(books, arrBook)  //copy方法中的第一个参数是目标数据，第二个参数是源数据

	//copy之后修改值，不会影响源切片的值
	books[0].author = "shiyifei"
	fmt.Println("after copying and modifing books[0], books:")
	printBookArr(books)

	fmt.Println("after copying and modifing books[0], arrBook:")
	printBookArr(arrBook)

	//切片的排序操作，按照name值大小升序排列
	sort.Slice(arrBook, func(i,j int) bool {return arrBook[i].name < arrBook[j].name})
	fmt.Println("after sort by book name, arrBook:")
	printBookArr(arrBook)

	sort.Slice(arrBook, func(i,j int) bool {return arrBook[i].price > arrBook[j].price})
	fmt.Println("after sort by book price desc, arrBook:")
	printBookArr(arrBook)
}

func printArr(arr []int) {
	fmt.Printf("length=%d,cap=%d,type is: %T, value is:%v\n", len(arr), cap(arr), arr, arr)
}

/*
	输出对象数组中的每一个元素
*/
func printBookArr(arr []Book) {
	for i := 0; i < len(arr); i++ {
		fmt.Printf("index=%d,name=%s,author=%s,pubtime=%s,price=%.2f\n", i, arr[i].name, arr[i].author, arr[i].pubtime, arr[i].price)
	}
}

func printBookArr1(arr []Book) {
	for k, v := range arr {
		fmt.Printf("index=%d,name=%s,author=%s,pubtime=%s,price=%.2f\n", k, v.name, v.author, v.pubtime, v.price)
	}
}
