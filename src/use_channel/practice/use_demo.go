/**
	本示例演示的是写入通道的值实际上是副本而不是引用，
	写入通道之后的值即使稍后有变化也不影响从通道中获取的值
	从通道拿出的其实也是副本
 */

package practice

import "fmt"

type Addr struct {
	city string
	district string
}
type Person struct {
	Name string
	Age int
	Address Addr
}

func TestCopy(){
	var personChan = make(chan Person, 1)

	person := Person{"wangzhongwei", 35, Addr{"Beijing", "haidian"}}
	fmt.Printf("person(1):%v \n", person)
	personChan <- person

	person.Address.district = "shijingshan"
	fmt.Printf("person(2):%v \n", person)

	personCopy := <-personChan
	fmt.Printf("person_copy:%v \n", personCopy)
}
