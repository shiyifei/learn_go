package practice

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	mrand "math/rand"
	"strings"
	"time"
)

//生成32位md5字串
func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

/**
 * 测试将reader中的数据读取到目标（byte切片)的方法
 */
func testReadFull() {
	r := strings.NewReader("some io.Reader stream to be read\n")
	readerLength := r.Len() //Len()返回未读的字符串字节数
	buf := make([]byte, 4)
	_, err := io.ReadFull(r, buf)
	if err != nil {
		fmt.Println("111, err:", err)
	}
	fmt.Printf("r.len=%d,buf:%s\n", r.Len(), buf) //这里的长度与初始长度不一样了，33-4=29

	longBuf := make([]byte, r.Len())
	_, err = io.ReadFull(r, longBuf)
	if err != nil {
		fmt.Println("222, err:", err)
	} else {
		fmt.Printf("longBuf:%s\n", longBuf)
	}

	currBuf := make([]byte, readerLength) //这里的长度是初始的字符串字节数
	_, err = io.ReadFull(r, currBuf)
	if err != nil {
		fmt.Println("333, err:", err)
	} else {
		fmt.Printf("currBuf:%s\n", currBuf)
	}

}

func GetRandomString(l int) string {
	str := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"
	bytes := []byte(str)
	result := []byte{}
	r := mrand.New(mrand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

//生成32位的Guid字串
func UniqueId() string {
	b := make([]byte, 48)

	_, err := io.ReadFull(rand.Reader, b)

	if err != nil {
		return ""
	}
	strA := base64.URLEncoding.EncodeToString(b)

	//strA是64位长度的字符串
	fmt.Printf("len of str:%d, value:[%s] \n", len(strA), strA)

	return GetMd5String(strA)
}

func TestGuid() {
	testReadFull()

	for i := 0; i < 10; i++ {
		output := UniqueId()
		fmt.Printf("len of str:%d, value:[%s] \n", len(output), output)
	}
}
