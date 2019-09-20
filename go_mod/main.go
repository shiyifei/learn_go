package main

import (
	    "log"

	        "github.com/toolkits/smtp"
	)

func main() {
	s := smtp.NewSMTP("smtp.163.com:25", "notverygood@163.com", "password",false,false,false)
	    log.Println(s.SendMail("notverygood@163.com", "625378510@qq.com;shi_yi_fei@163.com","这是subject", "这是body,<font color=red>red</font>"))
}
