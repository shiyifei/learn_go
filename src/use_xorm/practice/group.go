package practice

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var eg * xorm.EngineGroup

func Test() {
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
	eg, err = xorm.NewEngineGroup(master, slaves)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	eg.Ping()

}
