package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"use_file/writefile"
)

func main() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))  //获取当前路径
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("111, current directory:", dir)
	writefile.WriteStringToFile("hello,shiyf, what are you doing now ?")
	writefile.WriteByteToFile()
	strInputs := []string{"hello,shiyf,what are you doing now?","I am learning go language","I want to happy everyday"}
	writefile.WriteLineToFile(strInputs)
	var input string = "tomorrow is weekend"
	writefile.AppendStringToFile(input)
}
