package main

import (
	"bytes"
	"crypto"
	"encoding/hex"
	"fmt"
	"gopkg.in/gomail.v2"
	"hello/other"
	"strconv"
)


func sha1Encrypt(input string) string {
	buf := bytes.NewBufferString(input)
	h := crypto.SHA1.New()
	h.Write([]byte(buf.String()))
	output := hex.EncodeToString(h.Sum(nil))
	fmt.Printf("%s\n", output)
	return output
}

func main() {

	input := "user100000"
	output := sha1Encrypt(input)
	fmt.Println(input, output)
	return

	fmt.Println("Hello, shiyifei, what are you doing now ?")

	var a, b int
	a = 20
	b = 30

	area := other.GetArea(a, b)
	fmt.Println("area is:", area)

	var records []map[string]string
	var record map[string]string
	record = make(map[string]string)
	record["a"] = "areyouok"
	record["b"] = "howareyou"

	records = append(records, record)

	fmt.Println(record)

	fmt.Println(records)

	//定义收件人
	mailTo := []string{
		"shiyifei@xin.com",
		"625378510@qq.com",
		"notverygood@163.com",
	}
	//邮件主题为"Hello"
	subject := "Hello, this is an testing email111"
	// 邮件正文
	body := "Good,hello,shiyfei, what are you doing now? why can not send to you?"
	SendMail(mailTo, subject, body)

}

func SendMail(mailTo []string, subject string, body string) error {
	//定义邮箱服务器连接信息，如果是阿里邮箱 pass填密码，qq邮箱填授权码
	mailConn := map[string]string{
		"user": "notverygood@163.com",
		"pass": "SS971456jk",
		"host": "smtp.163.com",
		"port": "465",
	}

	port, _ := strconv.Atoi(mailConn["port"]) //转换端口类型为int

	m := gomail.NewMessage()
	m.SetHeader("From", "XD Game"+"<"+mailConn["user"]+">") //这种方式可以添加别名，即“XD Game”， 也可以直接用<code>m.SetHeader("From",mailConn["user"])</code> 读者可以自行实验下效果
	m.SetHeader("To", mailTo...)                            //发送给多个用户
	m.SetHeader("Subject", subject)                         //设置邮件主题
	m.SetBody("text/html", body)                            //设置邮件正文

	d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])

	err := d.DialAndSend(m)
	return err

}
