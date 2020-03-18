package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"use_xorm/practice"
)

var engine *xorm.Engine

func main() {
	engine, err := xorm.NewEngine("mysql", "manager:SYF!123mysql@tcp(192.168.56.102:3306)/test?charset=utf8")
	if err != nil {
		fmt.Println(err.Error())
	}
	//设置数据库连接池的最大打开连接数
	engine.SetMaxOpenConns(4)
	//result := engine.Ping()
	//fmt.Println(result)
	fmt.Println("are you ok?")
	practice.Test()
	practice.Method1()

}
