/**
自定义函数类型，可以作为其他函数的入参，也包括各种加密方法的实例包括sha,md5
*/
package practice

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"io"
)

//自定义的函数类型，函数值可以作为其他函数的参数
type Encipher func(input string) []byte

func UseType() {
	var input string
	input = "areyouok"
	var out string

	//函数类型的函数的调用方式一
	outFunc := GenEncryptionFunc(sha1Encrypt)
	out = outFunc(input)
	fmt.Println(out)

	//函数类型的函数的调用方式二，后面直接跟上入参
	out = GenEncryptionFunc(sha2Encrypt)(input)
	fmt.Println(out)

	input = "123456"
	out = GenEncryptionFunc(md5Encrypt)(input)
	fmt.Println(out)
}

/**
sha1加密 一种实现方式
*/
func sha1Encrypt(input string) []byte {
	h := sha1.New()
	io.WriteString(h, input)
	output := h.Sum(nil)
	fmt.Printf("output:")
	fmt.Println(output)
	return output
}

/**
sha1加密 另一种实现方式
*/
func sha2Encrypt(input string) []byte {
	h := sha1.New()
	h.Write([]byte(input))
	output := h.Sum(nil)
	fmt.Printf("output:")
	fmt.Println(output)
	return output
}

/**
  MD5加密字符串
*/
func md5Encrypt(input string) []byte {
	h := md5.New()
	h.Write([]byte(input))
	output := h.Sum(nil)
	fmt.Printf("output:")
	fmt.Println(output)
	return output
}

/**
该函数实际上为一个函数类型,
该函数不能直接返回值，调用时参考本例中的main()方法
其实相当于闭包，可以动态地调用加密方法
该函数的返回值是一个函数，所以需要针对其返回值再次调用才能得到结果
*/
func GenEncryptionFunc(encrypt Encipher) func(string) string {
	return func(input string) string {
		return fmt.Sprintf("%x", encrypt(input)) //将[]byte转换为string
	}
}
