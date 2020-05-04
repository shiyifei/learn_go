package practice
import (
	"fmt"
	"sort"
	"strings"
	)
type Book struct{
	name string 
	author string
	pubtime string
	price float32
}


type Role struct {
	Roleid int
	Parentid int
}

func TestSlice() {
	var roles map[int]Role
	roles = make(map[int]Role)
	role := Role{Roleid:270, Parentid:16}
	roles[270] = role
	role = Role{Roleid:274, Parentid:17}
	roles[274] = role
	role = Role{Roleid:282, Parentid:270}
	roles[282] = role
	role = Role{Roleid:283, Parentid:270}
	roles[283] = role
	role = Role{Roleid:285, Parentid:270}
	roles[285] = role
	role = Role{Roleid:286, Parentid:270}
	roles[286] = role
	role = Role{Roleid:292, Parentid:282}
	roles[292] = role
	role = Role{Roleid:293, Parentid:285}
	roles[293] = role

	fmt.Println("roles:", roles)

	//针对map的key进行排序,如果不排序的话可能最终结果会出错
	var keys []int
	for k := range roles {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	currSystemRoles := make([]int, 0)
	currSystemRoles = append(currSystemRoles, 270)

	//如果直接遍历roles,不能保证遍历的顺序是按照由小到大输出的。
	for _,k := range keys {
		fmt.Printf("v:%+v \n", roles[k])
		for _, v1 := range currSystemRoles {
			if roles[k].Parentid == v1 {
				fmt.Printf("k=%d, v.Parentid=%d,v1=%d \n",k, roles[k].Parentid, v1)
				currSystemRoles = append(currSystemRoles, k)
				break
			}
		}
	}

	fmt.Println("currSystemRoles:", currSystemRoles)

	//这里相当于 php的 implode(",", arrayInt)
	joinStr := strings.Replace(strings.Trim(fmt.Sprint(currSystemRoles), "[]")," ", ",",-1)
	fmt.Println("joinStr:", joinStr)
}


func BasicOperate() {
	//未设置长度的数组即为切片(slice)
	numbers := []int{1,2,4,5,6,7}
	printArr(numbers)

	//数组的赋值
	var arrInt = make([]int,3,10)
	arrInt[0] = 11
	arrInt[1] = 12
	arrInt[2] = 13
	printArr(arrInt)

	if  arrInt == nil {
		fmt.Println("arrInt is Empty")
	}
	//从数组中截取一段作为切片的值
	arrInt =  numbers[1:4]  //截取数组， 会包含下标为1,2,3的数据 4-1=3，长度应该为3
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
	fmt.Println(arrStr)  //数组可以直接输出，中间用空格分开，如果元素中有空格就会导致区分不出每一个元素。

	var book1 Book
	book1.name ="php"
	book1.author = "shiyf"
	book1.pubtime = "2019-04-21"
	book1.price = 36.5

	var book2 Book
	book2.name ="java"
	book2.author = "maliqun"
	book2.pubtime = "2018-01-01"
	book2.price =45.80

	var book3 Book
	book3.name="go"
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

	//对象切片
	var books = make([]Book, len(arrBook), 2*cap(arrBook))
	copy(books, arrBook)  //copy方法中的第一个参数是目标数据，第二个参数是源数据

	printBookArr(books)

}

func printArr(arr []int) {
	fmt.Printf("length=%d,cap=%d,type is: %T, value is:%v\n", len(arr), cap(arr), arr, arr)
}

/*
	输出对象数组中的每一个元素
 */
func printBookArr(arr []Book) {
	for i:=0; i<len(arr); i++ {
		fmt.Printf("name=%s,author=%s,pubtime=%s,price=%.2f\n",arr[i].name,arr[i].author,arr[i].pubtime,arr[i].price)
	}
}
