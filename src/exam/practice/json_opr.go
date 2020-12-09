package practice

import (
	"encoding/json"
	mapstructure "github.com/mitchellh/mapstructure"
	"fmt"
	"reflect"
)

type User struct {
	Id int `json:"user_id,string" mapstructure:"user_id"`		//这里注意要有mapstructure这个标签才行的
	UserName string `json:"username"`
	Age int8 `json:"age,string"`
	Email string `json:"email"`
}

type Response1 struct {
	Code int `json:"code"`
	Message string `json:"msg"`
	Data []User `json:"data"`
}


type Response2 struct {
	Code int `json:"code"`
	Message string `json:"msg"`
	Data []map[string]interface{} `json:"data"`
}

var strJson string

func init() {
	strJson = `{
				"code":0,
				"msg":"ok",
				"data":[{"user_id":"1", "username":"wangxiao", "age":"23", "email":"wangxiao@a.com"},
						{"user_id":"3", "username":"zhangle", "age":"28", "email":"zhanglei@1.com"}]
				}`
}

func JsonDecode() {
	//简单json串解析成结构体切片
	jsonDecode1()

	//先将json解析为map切片，然后再转换为结构体切片
	fmt.Println("===================================================")
	jsonDecode2()

	//灵活转换为结构体切片的一种写法
	var user User
	var userList = make([]User, 0)
	userList = jsonDecode3(user).Interface().([]User)
	fmt.Printf("userList is: %#v \n", userList)

	for _, v := range userList {
		fmt.Printf("row: %#v \n", v)
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
	将list中的数据由map[string]interface{}转换为User,这种方式很灵活
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

	var user User

	//PHP返回的json串解析时注意弱类型验证
	config := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		Result:           &user,
	}
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		panic(err)
	}

	//使用第三方类库将map[string]interface() 强制转换为 []User
	var objMap = make(map[string]interface{})
	for _,objMap = range obj.Data {
		fmt.Printf("data is: %#v \n", objMap)

		err = decoder.Decode(objMap)
		if err != nil {
			fmt.Println("err is:",err)
		} else {
			fmt.Printf("user is:%#v \n", user)
		}
	}
}

/**
	source 传入的一个结构体变量
	返回结构体切片
 */
func jsonDecode3(source interface{}) reflect.Value {
	var obj Response2
	err := json.Unmarshal([]byte(strJson), &obj)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Printf("obj is:%+v \n", obj)
	//
	//fmt.Printf("obj.Code is: %+v \n", obj.Code)
	//fmt.Printf("obj.Message is: %+v \n", obj.Message)


	sourceType := reflect.TypeOf(source)
	//kind := reflect.Kind(source)

	//fmt.Println(sourceType)

	//PHP返回的json串解析时注意弱类型验证
	config := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		Result:           &source,
	}
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		panic(err)
	}

	//使用第三方类库将map[string]interface() 强制转换为 []User
	var objMap = make(map[string]interface{})

	sliceType := reflect.SliceOf(sourceType)
	retArr := reflect.MakeSlice(sliceType, 0, 0)

	//fmt.Printf("retArr is: %#v \n", retArr)

	for _,objMap = range obj.Data {
		fmt.Printf("data is: %#v \n", objMap)

		err = decoder.Decode(objMap)
		if err != nil {
			fmt.Println("err is:",err)
		} else {
			retArr = reflect.Append(retArr, reflect.ValueOf(source))
			//fmt.Printf("source is:%#v \n", source)
		}
	}
	//fmt.Printf("retArr is:%#v \n", retArr)
	return retArr
}

