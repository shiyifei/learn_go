package practice

import (
	"encoding/json"
	mapstructure "github.com/mitchellh/mapstructure"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Order struct {
	Name        string   `json:"name"`
	OrderItem   []Item   `json:"item"`
	OrderRefund []Refund `json:"refund"`
}

type Item struct {
	Name string `json:"name"`
	Oid  int    `json:"oid"`
}

type Refund struct {
	Name    string `json:"name"`
	Item    int    `json:"item"`
	Content string `json:"content"`
	Imgs    string `json:"imgs"`
	Status  string `json:"status"`
}

type Message struct {
	Id   int
	Name string
}

type Response1 struct {
	Code int `json:"code"`
	Message string `json:"msg"`
	Data []User `json:"data"`
 }
type User struct {
	UserId int `json:"user_id" mapstructure:"user_id"`		//这里注意要有mapstructure这个标签才行的
	Username string `json:"username"`
	Age byte `json:"age"`
	Email string `json:"email"`
}

type Response2 struct {
	Code int `json:"code"`
	Message string `json:"msg"`
	Data []map[string]interface{} `json:"data"`
}

var strJson string

func init() {
	strJson = `{
				"code":200,
				"msg":"ok",
				"data":[{"user_id":1, "username":"wangxiao", "age":23, "email":"wangxiao@a.com"},
						{"user_id":3, "username":"zhangle", "age":28, "email":"zhanglei@1.com"}]
				}`
}

func UseJson() {
	var obj = ajaxReturn("ok", "")
	fmt.Printf("obj.code=%s,obj.msg=%s \n", obj["code"], obj["msg"])
	jsonToStruct()
	structToJson()
}


func JsonEncodeDecode() {
	fmt.Println("arrive in here ,JsonEncodeDecode() ")
	jsonDecode1()
	fmt.Println("===================================================")
	jsonDecode2()
}

func ajaxReturn(code, msg string) map[string]string {
	return map[string]string{
		"code": code,
		"msg":  msg,
	}
}

/**
	直接转成预定类型的变量
 */
func jsonDecode1() {
	var obj Response1
	err := json.Unmarshal([]byte(strJson), &obj)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("obj is:%+v \n", obj)

	fmt.Printf("obj.Code is: %+v \n", obj.Code)
	fmt.Printf("obj.Message is: %+v \n", obj.Message)
	for _,v := range obj.Data {
		fmt.Printf("data is: %#v \n", v)
	}
}

/**
	将list中的数据由map[string]interface{}转换为User
 */
func jsonDecode2() {
	var obj Response2
	err := json.Unmarshal([]byte(strJson), &obj)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("obj is:%+v \n", obj)

	fmt.Printf("obj.Code is: %+v \n", obj.Code)
	fmt.Printf("obj.Message is: %+v \n", obj.Message)

	//使用第三方类库将map[string]interface() 强制转换为 []User
	var mapInstance = make(map[string]interface{})
	for _,mapInstance = range obj.Data {
		fmt.Printf("data is: %#v \n", mapInstance)
		//mapInstance = v
		var user User
		err = mapstructure.Decode(mapInstance, &user)
		if err != nil {
			fmt.Println("err is:",err)
		} else {
			fmt.Printf("user is:%#v \n", user)
		}
	}
}





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

/*func json_decode(json string) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	err := json.Unmarshal(json)
	if err != nil {
		return nil, err
	}
	return result, nil
}*/
