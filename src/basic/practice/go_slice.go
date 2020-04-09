package practice
import "fmt"
type Book struct{
	name string 
	author string
	pubtime string
	price float32
}


func Go_slice() {
	numbers := []int{1,2,4,5,6,7}

	printArr(numbers)
	
	var arrInt = make([]int,3,10)
	arrInt[0] = 11
	arrInt[1] = 12
	arrInt[2] = 13

	printArr(arrInt)

	var arrStr [3]string
	arrStr[0] = "are you ok ?"
	arrStr[1] = "hello"
	arrStr[2] = "how do you do"

	fmt.Println(arrStr)  //数组可以直接输出，中间用空格分开，如果元素中有空格就会导致区分不出每一个元素。
	
	
	if  arrInt == nil {
		fmt.Println("arrInt is Empty")
	}
	arrInt =  numbers[1:4]  //截取数组， 会包含下标为1,2,3的数据 4-1=3，长度应该为3
	fmt.Printf("arrInt,length=%d,cap=%d,value=%v\n",len(arrInt),cap(arrInt),arrInt)

	arrObj := numbers[0:5]
	printArr(arrObj)

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

/**
输出对象数组中的每一个元素
 */
func printBookArr(arr []Book) {
	for i:=0; i<len(arr); i++ {
		fmt.Printf("name=%s,author=%s,pubtime=%s,price=%.2f\n",arr[i].name,arr[i].author,arr[i].pubtime,arr[i].price)
	}
}
