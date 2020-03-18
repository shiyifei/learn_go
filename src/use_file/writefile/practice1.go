package writefile

import (
	"fmt"
	"os"
)

/**
	写入字符串到文件中
 */
func WriteStringToFile(input string) {
	//  /var/www/html/learn_go/src/use_file/writefile/test.txt
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := f.WriteString(input)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

/**
	写入字节切片到文件中
 */
func WriteByteToFile() {
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	byteArr := []byte{104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100}
	l, err := f.Write(byteArr)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "bytes written succefully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

/**
	将字符串切片一行一行地写入到文件中
 */
func WriteLineToFile(strInputs []string) {
	f, err := os.Create("file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _,v := range strInputs {
		_, err = fmt.Fprintln(f, v)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("file written successfully")
}

/**
	向文件中追加写入内容
 */
func AppendStringToFile(input string) {
	f, err := os.OpenFile("file.txt", os.O_APPEND|os.O_WRONLY, 0644)  //以追加和写的方式打开文件
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = fmt.Fprintln(f, input)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("file appended successfully")
}