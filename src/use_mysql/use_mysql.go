/**
	golang操作数据库的实例，包括增删改查
 */
package main

import (
	"fmt"
	"use_mysql/practice"
)

func main() {
	insertId, err := practice.Insert("lixiaoyun", "lixiaoyun@music.com")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(insertId)

	num, err := practice.Update(200036)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Println(num)

	num, err = practice.Delete(200034)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(num)

	output := practice.ReadData()
	fmt.Println(output)
}