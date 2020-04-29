package dao

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var DB * xorm.EngineGroup

func init() {
	var err error
	master, err := xorm.NewEngine("mysql", "manager:SYF!123mysql@tcp(192.168.56.102:3306)/test?charset=utf8")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//fmt.Println("master ping()", master.Ping())
	slaver, err := xorm.NewEngine("mysql", "manager:SYF!123mysql@tcp(192.168.56.102:3307)/test?charset=utf8")

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//fmt.Println("slaver ping()", slaver.Ping())

	slaves := []*xorm.Engine{slaver}
	DB, err = xorm.NewEngineGroup(master, slaves)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	DB.Ping()

}

