package practice

import (
	"fmt"
	"time"
)

func UseTime() {
	template := "2006-01-02 15:04:05"
	a := time.Now()
	b := time.Now().Local()
	c := time.Now().Format(template)
	fmt.Println(a, "\t", b, "\t", c)

	template = "2006010215"

	output := time.Now().Format(template)
	fmt.Println("output:", output)
}
