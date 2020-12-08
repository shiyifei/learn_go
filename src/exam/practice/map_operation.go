package practice

import (
	"encoding/json"
	"fmt"
)

type MapUser map[int]User

type MyString string
var userList MapUser


func (s MyString) add() int {
	var it = 0
	fmt.Println(it)
	return it
}

func init() {
	//userList = make(MapUser)
	//mapA := make(MapUser)
	//userList = &mapA

	//userList = new (MapUser)

	var maps map[int]User
	//aa := new(map[int]User)
	maps = make(map[int]User)
	//maps = *aa





	var user User
	user.Id = 1
	user.UserName = "liubei"
	user.Age = 35
	user.Email = "liubei@1.com"
	maps[user.Id] = user

	user.Id = 2
	user.UserName = "zhangfei"
	user.Age = 34
	user.Email = "zhangfei@1.com"
	maps[user.Id] = user

	fmt.Println("maps:",maps)
	return
}

/**
	增加一个用户
 */
func (users *MapUser) AddOneUser(user User) bool {
	if !users.userIsExist(user) {
		(*users)[user.Id] = user
		return true
	} else {
		fmt.Println("username:"+user.UserName+" is exist")
		return false
	}
}

/**
删除一个用户
 */
func (users *MapUser) deleteOneUser(user User) {
	_, ok := (*users)[user.Id]
	if ok == true {
		delete(*users, user.Id)
	} else {
		fmt.Println("user id:",user.Id," is not exist")
	}
}
/**
  修改一个用户
 */
func (users *MapUser) modifyOneUser(user User) {
	_, isExist := (*users)[user.Id]
	if isExist {
		(*users)[user.Id] = user
	} else {
		fmt.Println("user:",user.Id," is not exist")
	}
}

/**
	根据用户名查找一个用户
 */
func (users *MapUser) userIsExist(user User) bool {
	for _, v := range *users {
		if v.UserName == user.UserName {
			return true
		}
	}
	return false
}

/**
	打印所有用户
 */
func (users *MapUser) printUsers(){
	bytes, _ := json.Marshal(*users)
	fmt.Printf("%s \n", bytes)
}

/**
整体测试方法
 */
func MapOperation() {
	//defer语句会保证操作完毕才输出
	defer userList.printUsers()

	user3 := User{Id:3, UserName:"guanyu", Age:33, Email:"guanyu@1.com"}
	user4 := User{Id:4, UserName:"caocao", Age:36, Email:"caocao@1.com"}
	user5 := User{Id:5, UserName:"sunquan", Age:33, Email:"sunquan@1.com"}

	userList.AddOneUser(user3)
	userList.AddOneUser(user4)
	userList.AddOneUser(user5)

	//添加一条重复的记录，看是否有回应
	user6 := User{Id:6, UserName:"guanyu", Age:32, Email:"guanyu@2.com"}
	userList.AddOneUser(user6)

	user3.Email = "guanyu@sina.com"
	userList.modifyOneUser(user3)

	user4.Age = 42
	userList.modifyOneUser(user4)

	userList.deleteOneUser(user5)
}
