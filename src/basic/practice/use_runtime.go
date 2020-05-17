package practice

//runtime.Caller函数原型： func Caller(skip int) (pc uintptr, file string, line int, ok bool)

import (
	"errors"
	"fmt"
	"runtime"
)

func Use_runtime() {

	currentFile, line := currentFileAndLine()
	fmt.Printf("current file:%s, line number is:%d \n", currentFile, line )
}

//获取当前文件名和行数的方法
func currentFileAndLine() (string, int){
	//Caller()方法中的skip参数为1表示currentFileAndLine方法的调用位置
	//skip参数为0,表示当前行
	_, file, line, ok := runtime.Caller(0)
	if !ok {
		panic(errors.New("Can not get current file and line number"))
	}
	return file, line
}
