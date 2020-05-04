package practice

import (
	"github.com/gohouse/converter"
	"fmt"
)


func Table2Struct(tableName string)
{
	err := converter.NewTable2Struct()
	.SavePath("/var/www/html/learn_go/bin/model.go")
	.Dsn("root:SYF!123mysql@tcp(192.168.1.102:3306)/test?charset=utf8")
	.Run()
	if err != nil {
		fmt.Println("err:", err )
	}
}
