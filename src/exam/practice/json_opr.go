package practice

import (
	"encoding/json"
	mapstructure "github.com/mitchellh/mapstructure"
	"fmt"
	"reflect"
	"time"
)

type User struct {
	Id int `json:"user_id,string" mapstructure:"user_id"`		//这里注意要有mapstructure这个标签才行的
	UserName string `json:"username"`
	Age int8 `json:"age,string"`				//注意这里的age 需要加入string,不然只能json串中age为整数的值
	Email string `json:"email"`
	Gender int8 `json:"gender,string"`	//如果json串中没有该值会获取到结果，转换后会有一个默认值。
	Mobile string						//如果没有写json Tag，如果json串中有该变量的小写字母key对应的值，也是可以的。
}

type Employee struct {
	EmpId int `json:"user_id,string" mapstructure:"user_id"`		//这里注意要有mapstructure这个标签才行的
	UserName string `json:"username"`
	Age int8 `json:"age,string"`
	Email string `json:"email"`
	Gender int8 `json:"gender,string"`
	Mobile string
}

type ResponseUser struct {
	Code int `json:"code"`
	Message string `json:"msg"`
	Data struct {
		Total int `json:"total"`
		Page int `json:"page"`
		List []User `json:"list"`
	} `json:"data"`
}

type Data struct {
	Total int `json:"total"`
	Page int `json:"page"`
	List []interface{} `json:"list"`
}

type Response struct {
	Code int `json:"code"`
	Message string `json:"msg"`
	Data Data `json:"data"`
}

type ResponseMap struct {
	Code int `json:"code"`
	Message string `json:"msg"`
	Data struct {
		Total int `json:"total"`
		Page int `json:"page"`
		List []map[string]interface{} `json:"list"`
	} `json:"data"`
}

var strJson string

func init() {
	strJson = `{
				"code":0,
				"msg":"ok",
				"data":{
						"total":2,
						"page":1,
						"list":[
						{"user_id":"1", "username":"wangxiao", "age":23, "email":"wangxiao@a.com","mobile":"13712223333"},
						{"user_id":"3", "username":"zhangle", "age":"28", "email":"zhanglei@1.com","mobile":"13712223333"}
						]
                       }
				}`
}

func JsonDecode() {
	begin := time.Now()
	//简单json串解析成结构体切片
	directJsonDecode()
	diff :=  time.Now().Sub(begin)
	fmt.Println("directJsonDecode()，耗时：",diff)

	fmt.Println("=======================end directJsonDecode()===========================================")

	//先将json解析为map切片，然后再转换为[]interface{}

	begin = time.Now()
	//传入User类型变量，最终返回结果中的列表类型实际上是[]User
	var user User
	ret := commonJsonDecode(user)
	diff =  time.Now().Sub(begin)
	fmt.Println("commonJsonDecode()，耗时：",diff)

	fmt.Printf("ret is:%#v \n", ret)
	for _, item := range ret.Data.List {
		user, ok := item.(User)	//interface{} 强转为 User
		if ok {
			fmt.Printf("user is:%#v\n", user)
		}
	}
	fmt.Println("========================================================================================")
	//传入Employee类型变量，最终返回结果中的列表类型实际上是[]Employee
	var emp Employee
	ret = commonJsonDecode(emp)
	fmt.Printf("ret is:%#v \n", ret)

	for _, item := range ret.Data.List {
		user, ok := item.(Employee)
		if ok {
			fmt.Printf("user is:%#v\n", user)
		}
	}

	fmt.Println("=======================end commonJsonDecode()===========================================")
	/*
	//灵活转换为结构体切片的另一种写法
	var userList = make([]User, 0)
	begin = time.Now()
	userList = reflectJsonDecode(user).Interface().([]User)
	diff =  time.Now().Sub(begin)
	fmt.Println("commonJsonDecode()，耗时：",diff)
	fmt.Printf("userList is: %#v \n", userList)

	for _, v := range userList {
		fmt.Printf("row: %#v \n", v)
	}
	fmt.Println("=======================end ReflectJsonDecode()===========================================")
	*/

}

/**
	直接转成预定类型的变量
 */
func directJsonDecode() {
	var obj ResponseUser
	err := json.Unmarshal([]byte(strJson), &obj)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("obj is:%#v \n", obj)
	userList := obj.Data.List
	for _,v := range userList {
		fmt.Printf("user is: %#v \n", v)
	}
}

/**
 将list中的数据由map[string]interface{}转换为[]User,这种方式很灵活
 */
func commonJsonDecode(target interface{}) Response {
	var obj ResponseMap
	err := json.Unmarshal([]byte(strJson), &obj)
	if err != nil {
		fmt.Println(err)
	}
	//PHP返回的json串解析时注意弱类型验证
	config := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		Result:           &target,
	}
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		panic(err)
	}

	//使用第三方类库将map[string]interface() 强制转换为 []User
	var objMap = make(map[string]interface{})

	var objData Response
	var userArr = make([]interface{}, 0)
	objData.Code = obj.Code
	objData.Message = obj.Message
	var innerData Data
	innerData.Page = obj.Data.Page
	innerData.Total = obj.Data.Total
	for _,objMap = range obj.Data.List {
		fmt.Printf("origin row is: %#v \n", objMap)
		err = decoder.Decode(objMap)
		if err != nil {
			fmt.Println("err is:",err)
		} else {
			fmt.Printf("converted struct is:%#v \n", target)
			userArr = append(userArr, target)
		}
	}
	innerData.List = userArr
	objData.Data = innerData
	return objData
}

/**
	source 传入的一个结构体变量, 返回的是反射类型的值（实际上是结构体切片)
 */
func reflectJsonDecode(target interface{}) reflect.Value {
	var obj ResponseMap
	err := json.Unmarshal([]byte(strJson), &obj)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("obj is:%#v \n", obj)
	data  := obj.Data
	sourceType := reflect.TypeOf(target)
	//kind := reflect.Kind(target)

	//fmt.Println(sourceType)

	//PHP返回的json串解析时注意弱类型验证
	config := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		Result:           &target,
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

	for _,objMap = range data.List {
		fmt.Printf("origin data is: %#v \n", objMap)

		err = decoder.Decode(objMap)
		if err != nil {
			fmt.Println("err is:",err)
		} else {
			retArr = reflect.Append(retArr, reflect.ValueOf(target))
			//fmt.Printf("target is:%#v \n", target)
		}
	}
	//fmt.Printf("retArr is:%#v \n", retArr)
	return retArr
}

