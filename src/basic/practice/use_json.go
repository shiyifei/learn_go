package practice

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Order struct {
	Name        string   `json:name`
	OrderItem   []Item   `json:item`
	OrderRefund []Refund `json:refund`
}

type Item struct {
	Name string `json:name`
	Oid  int    `json:oid`
}

type Refund struct {
	Name    string `json:name`
	Item    int    `json:item`
	Content string `json:content`
	Imgs    string `json:imgs`
	Status  string `json:status`
}

type Message struct {
	Id   int
	Name string
}

type Response struct {
	Code int `json:code`
	Message string `json:msg`
	Data []interface{} `json:data`
 }
type User struct {
	Id int `json:user_id`
	Username string `json:username`
	Age byte `json:age`
	Email string `json:email`
}

var strJson String

func init() {
	strJson = `{
				"code":0,
				"msg":"ok",
				"data":[{"user_id":1, "username":"wangxiao", "age":23, "email":"wangxiao@a.com"},
						{"user_id":3, "username":"zhangle", "age":28, "email":"zhanglei@1.com"},
						]
				}`
}

func UseJson() {
	var obj = ajaxReturn("ok", "")
	fmt.Printf("obj.code=%s,obj.msg=%s \n", obj["code"], obj["msg"])
	jsonToStruct()
	structToJson()
}

func ajaxReturn(code, msg string) map[string]string {
	return map[string]string{
		"code": code,
		"msg":  msg,
	}
}

/*
func jsonDecode() {
	dec := json.NewDecoder(strings.NewReader(strJson))
	for {
		var obj Response
		err := dec.Decode(&obj)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("%d:%s \n", obj.Code, obj.Message)
			for k, v := range obj.Data {
				//todo:如何将[]interface{}转换为[]User
			}
		}
}
*/


//json转结构体
func jsonToStruct() {
	const jsonStream = `
			{"Id":11, "Name":"a"}
			{"Id":12, "Name":"b"}
			{"Id":13, "Name":"c"}
			{"Id":14, "Name":"d"}
			{"Id":15, "Name":"e"}	
		`
	dec := json.NewDecoder(strings.NewReader(jsonStream))
	for {
		var m Message
		err := dec.Decode(&m)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("%d:%s \n", m.Id, m.Name)
		}
	}
}

/**
结构体转json
*/
func structToJson() {
	var m Order
	m.Name = "order10010"
	m.OrderItem = append(m.OrderItem, Item{Name: "shanghai_house", Oid: 1})
	m.OrderItem = append(m.OrderItem, Item{Name: "beijing_car", Oid: 2})
	for i := 1; i < 6; i++ {
		str := []byte("Things")
		str = strconv.AppendInt(str, int64(i), 10)
		orderi := Item{Name: string(str), Oid: i}
		m.OrderItem = append(m.OrderItem, orderi)
	}
	bytes, _ := json.Marshal(m)
	fmt.Printf("json.m,%s \n", bytes)
}

func json_encode(data interface{}) (string, error) {
	bytes, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func json_decode(json string) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	err := json.Unmarshal(json)
	if err != nil {
		return nil, err
	}
	return result, nil
}
