package service

import (
	"fmt"
	"reflect"
	"time"
	"use_xorm/dao"
	"github.com/gogf/gf/util/gconv"
)

type Users struct{
	Id int
	UserName string
	Email string
	CreateTime time.Time
}

type UserService struct {

}


func (p *UserService) GetUserList1(params map[string]interface{}) []Users{

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


func (p *UserService) GetUserList2(params map[string]interface{}) []Users{

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