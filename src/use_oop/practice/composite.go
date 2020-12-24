/**
 * 以下演示面向对象中的组合模式的使用示例
 * @Author:shiyf
 * @Date: 2020/12/25 0:01
 **/
/**
	这是Golang的组合模式，可以实现OOP的继承。
	被组合的类型People所包含的方法虽然升级成了外部类型Teacher这个组合类型的方法（一定要是匿名字段），
    但它们的方法(ShowA())调用时接收者并没有发生变化。
	此时People类型并不知道自己会被什么类型组合，当然也就无法在调用方法时去使用未知的组合者Teacher类型的功能。
 */
package practice

import "fmt"

type people struct{}

func (p *people) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}

func (p *people) ShowB() {
	fmt.Println("showB")
}

type teacher struct {
	people
}

func (t *teacher) ShowB() {
	fmt.Println("teacher showB")
}

func Test() {
	ter := teacher{}
	ter.ShowA()
	ter.ShowB()
	fmt.Println("============================")
}
