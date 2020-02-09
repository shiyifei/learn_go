package practice
import "fmt"

func Fibnach() {
	fmt.Println(fib(34))
}

func fib(n int64) int64 {
	if n<2 {
		return n
	}
	return fib(n-2) + fib(n-1)
}

