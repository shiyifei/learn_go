<<<<<<< HEAD
package main

import(
	UseMysql "use_mysql/practice"
)

func main() {
	dbAccess := UseMysql.ConnectToDB()
	dbAccess.InsertOne()
	dbAccess.SelectMultiUser()
	dbAccess.SelectOneUser(2001)
	dbAccess.UpdateRecord(2001)
	dbAccess.DeleteRecord(2001)
	dbAccess.ExecTrans()
}
=======
/**
	golang操作数据库的实例，包括增删改查
 */
package main

import (
	"fmt"
	"use_mysql/practice"
)

func main() {
	insertId, err := practice.Insert("lixiaoyun", "lixiaoyun@music.com")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(insertId)

	num, err := practice.Update(200036)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Println(num)

	num, err = practice.Delete(200034)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(num)

	output := practice.ReadData()
	fmt.Println(output)
}
>>>>>>> a8c918e90ac59047606ba621d6d950916389596e
