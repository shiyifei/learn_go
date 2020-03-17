package main

import (
	    "log"
	    "github.com/toolkits/smtp"
	    UseMysql "go_mod/practice"
	)

func main() {
	s := smtp.NewSMTP("smtp.163.com:25", "notverygood@163.com", "SS971456jk",false,false,false)
	log.Println(s.SendMail("notverygood@163.com", "625378510@qq.com;shi_yi_fei@163.com","这是subject", "这是body,<font color=red>red</font>"))

	//测试存取mysql数据
	dbAccess := UseMysql.ConnectToDB()
	dbAccess.InsertOne()
	dbAccess.SelectMultiUser()
	dbAccess.SelectOneUser(2001)
	dbAccess.UpdateRecord(2001)
	dbAccess.DeleteRecord(2001)
	dbAccess.ExecTrans()
}
