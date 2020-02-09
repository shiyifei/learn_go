package practice

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			//x, y = y, x+y

			temp := x
			x = y
			y = temp + y

			case <-quit:
				fmt.Println("quit")
				return
			/*default:
			fmt.Println("no value received")*/
		}
	}
}

func AboutSelect() {
	c := make(chan int)
	quit := make(chan int)

	defer fibonacci(c, quit)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

}
