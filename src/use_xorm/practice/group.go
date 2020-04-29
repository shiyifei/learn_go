package practice

import (
	"fmt"
	"use_xorm/service"
)

func GetUserList() {

	 service := &service.UserService{}
	 params := map[string]interface{}{"id":20000, "limit":3}
	 ret := service.GetUserList1(params)
	 fmt.Printf("%+v \n", ret)

	 result := service.GetUserList2(params)
	fmt.Printf("%+v\n", result)
}
