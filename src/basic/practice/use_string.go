package practice

import (
	"fmt"
	"unicode/utf8"
)

func UseString() {
	var name string
	name= "hello, what are you doing now?"
	printBytes(name)
	printChars(name)
	fmt.Println(changeStr([]rune(name)))


	name= "all say：天行健，君子以自强不息"
	printBytes(name)
	printChars(name)
	printCharsAndBytes(name)

	//获取字符串实际文字长度
	length := getStrLength(name)
	fmt.Printf("length of str:[%s] is %d \n", name, length)
}


func printBytes(str string) {
	for i:=0; i<len(str); i++ {
		fmt.Printf("%x ", str[i])
	}
	fmt.Println()
}

/**
如果需要打印多字节字符（比如中文字符（utf8编码)就占用了三个字节）
 */
func printChars(str string) {
	fmt.Println("length of str:", len(str))
	runes := [] rune(str)  //这里字符串被转换为一个rune切片, rune是32位有符号整形 int32的别名
	fmt.Println("length of runes:", len(runes))
	for i:=0; i< len(runes); i++ {
		fmt.Printf("%c ", runes[i])
	}
	fmt.Println()
}

/**
	使用for range 遍历字符串
 */
func printCharsAndBytes(str string) {
	for index,rune := range str {
		fmt.Printf("%c starts at byte %d \n", rune, index)
	}
}
/**
	获取字符串文字长度
 */
func getStrLength(str string) int {
	return utf8.RuneCountInString(str)
}

/**
	如何修改字符串
 */
func changeStr(runes []rune) string {
	runes[0] = 'H'
	return string(runes)
}


