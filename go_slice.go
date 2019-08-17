package main
import "fmt"
type Book struct{
	name string 
	author string
	pubtime string
	price float32
}


func main() {
	numbers := []int{1,2,4,5,6,7}
	
	var arrInt = make([]int,3,10)
	arrInt[0] = 11
	arrInt[1] = 12
	arrInt[2] = 13

	printArr(arrInt)
	
	
	if(arrInt == nil) {
		fmt.Println("arrInt is Empty")
	}
	arrInt =  numbers[1:4]

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

	var arrBook []Book
	arrBook = append(arrBook, book1)
	arrBook = append(arrBook, book2)
	arrBook = append(arrBook, book3)
	printBookArr(arrBook)

	var books=make([]Book,len(arrBook),2*cap(arrBook))
	copy(books,arrBook)

	printBookArr(books)

}
func printArr(arr []int) {
	fmt.Printf("length=%d,cap=%d,value is:%v\n", len(arr),cap(arr),arr)
}

func printBookArr(arr []Book) {
	for i:=0;i<len(arr);i++ {
		fmt.Printf("name=%s,author=%s,pubtime=%s,price=%.2f\n",arr[i].name,arr[i].author,arr[i].pubtime,arr[i].price)
	}
}
