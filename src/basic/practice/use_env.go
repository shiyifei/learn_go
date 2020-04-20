/**
 * os.Getenv检索环境变量并返回值，如果变量是不存在的，获取的值将是空的。
 */
package practice

import (
	"fmt"
	"os"
)

func UseEnv() {
	var1 := os.Getenv("GOPATH")
	fmt.Println("var1:", var1)

	var2 := os.Getenv("ACTIVE")
	fmt.Println("var2:", var2)
}
