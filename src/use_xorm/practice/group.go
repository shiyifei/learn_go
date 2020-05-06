package practice

import (
	"fmt"
	"use_xorm/service"
)

func GetUserList() {

	service := &service.UserService{}
	params := map[string]interface{}{"id": 20000, "limit": 3}
	ret := service.GetUserList1(params)
	fmt.Printf("%+v \n", ret)

	result := service.GetUserList2(params)
	fmt.Printf("%+v\n", result)
}

func GetUserInfo() {
	service := &service.UserService{}
	params := map[string]interface{}{"email": "caihewei@yonyou.com"}
	retA := service.GetUserByEmail(params)
	fmt.Printf("%+v \n", retA)
}

func TestFind() {
	service := &service.UserService{}
	params := map[string]interface{}{"id": 20000, "limit": 5}
	ret := service.TestFind(params)
	fmt.Printf("%+v \n", ret)
}

func TestRows() {
	service := &service.UserService{}
	params := map[string]interface{}{"id": 20000, "limit": 5}
	ret := service.TestRows(params)
	fmt.Printf("in TestRows(), ret:%+v \n", ret)
}
