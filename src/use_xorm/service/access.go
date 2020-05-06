package service

import (
	"fmt"
	"github.com/gogf/gf/util/gconv"
	"reflect"
	// "time"
	"use_xorm/dao"
)

type Users struct {
	Id       int    `json:"id" xorm:"not null pk autoincr INT(10) 'id'"`
	Username string `json:"username" xorm:"not null default '' VARCHAR(20) 'username'"`
	Email    string `json:"email" xorm:"not null default '' VARCHAR(30) 'email'"`
}

type UserService struct {
}

func (p *UserService) GetUserByEmail(params map[string]interface{}) []Users {
	var cols = []string{"id", "username", "createtime"}
	var valuesMap = make([]interface{}, len(cols))
	has, err := dao.DB.Table("users").Where("email=?", params["email"]).Cols(cols...).Get(&valuesMap)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Printf("has:%s, %v \n", reflect.TypeOf(has), has)
	users := make([]Users, 0)
	gconv.Struct(valuesMap, users)
	return users
}

func (p *UserService) GetUserList1(params map[string]interface{}) []Users {

	sql := "select * from users where id>? order by id desc limit 0,?"
	result, err := dao.DB.SQL(sql, params["id"], params["limit"]).QueryInterface()
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Printf("%+v \n", result)
	fmt.Printf("%+v \n", result[0])
	rettype := reflect.TypeOf(result[0]["email"])
	fmt.Println("email:", result[0]["email"])
	fmt.Printf("type=%s,email=%s \n", rettype, result[0]["email"])

	sql = "select count(id) as num from users where id>?"
	ret, err := dao.DB.SQL(sql, params["id"]).QueryInterface()
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Printf("%+v \n", ret)
	fmt.Printf("%+v \n", ret[0])
	rettype = reflect.TypeOf(ret[0]["num"])
	fmt.Printf("type=%s,count=%d \n", rettype, ret[0]["num"])
	var list []Users
	list = make([]Users, 0)
	user := &Users{}
	for _, v := range result {
		err := gconv.Struct(v, user)
		if err != nil {
			fmt.Println("err:", err)
		}
		list = append(list, *user)
	}
	return list
}

func (p *UserService) GetUserList2(params map[string]interface{}) []Users {

	sql := "select * from users where id>? order by id desc limit 0,?"
	result, err := dao.DB.SQL(sql, params["id"], params["limit"]).QueryString()
	if err != nil {
		fmt.Println("err:", err)
	}
	//return result
	//fmt.Println(result)
	var list []Users
	list = make([]Users, 0)
	user := &Users{}
	for _, v := range result {
		err := gconv.Struct(v, user)
		if err != nil {
			fmt.Println("err:", err)
		}
		list = append(list, *user)
	}
	return list
}

func (p *UserService) TestFind(params map[string]interface{}) map[int64]Users {
	userArr := make(map[int64]Users)
	// userArr := make([]Users, 0)
	var fields = []string{"id", "username", "email"}
	userArr[123] = Users{Id: 123, Username: "shiyifei", Email: "areyouok@163.com"}
	err := dao.DB.Table("users").Where("id>?", params["id"]).Cols(fields...).Limit(params["limit"].(int)).Find(&userArr)
	if err != nil {
		fmt.Println("in access.go, err:", err)
	}
	fmt.Printf("in access.go, %+v \n", userArr)
	return userArr
}
