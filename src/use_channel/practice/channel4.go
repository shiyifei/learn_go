package practice

import (
	"fmt"
	"time"
)

func channelNum(number int, chOutput chan int) {
	for {
		digit := number % 10
		chOutput <- digit
		number /= 10
		if number == 0 {
			break
		}
	}
	close(chOutput)
}

func calcuSquare(number int, chSquare chan int) {
	chOut := make(chan int)
	go channelNum(number, chOut)
	sum := 0
	for num := range chOut {
		sum += num * num
	}
	chSquare <- sum
}

func calcuCube(number int, chCube chan int) {
	chanUsed := make(chan int)
	go channelNum(number, chanUsed)
	sum := 0
	for num := range chanUsed {
		sum += num * num * num
	}
	chCube <- sum
}

func ShowResult(input int) {
	chanSquare := make(chan int)
	chanCube := make(chan int)
	go calcuSquare(input, chanSquare)
	go calcuCube(input, chanCube)
	square, cube := <-chanSquare, <-chanCube
	fmt.Println("in ShowResult(), output=",square+cube)
	time.Sleep(200*time.Millisecond)
}
