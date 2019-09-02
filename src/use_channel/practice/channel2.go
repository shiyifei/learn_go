/**
	关于信道的第二个示例， channel是如何在goroutine中使用的
 */
package practice

import "fmt"

/**
	计算number各位上的数字的平方和
 */
func calcuSquares(number int, chSquar chan int) {
	sum := 0
	for number != 0 {
		num := number % 10
		sum += num * num
		number /= 10
	}
	chSquar <- sum
}

/**
	计算number各位上的数字的立方和
 */
func calcuCubes(number int, chCube chan int) {
	sum := 0
	for number !=0 {
		num := number % 10
		sum += num * num * num
		number /= 10
	}
	chCube <- sum
}

/**
	打印最终计算的结果
 */
func PrintResult(input int) {
	chSquar := make(chan int)
	chCube := make(chan int)

	go calcuSquares(input, chSquar)
	go calcuCubes(input, chCube)

	square, cube := <-chSquar, <-chCube

	fmt.Println("in PrintResult(),result is :", square+cube)

}
