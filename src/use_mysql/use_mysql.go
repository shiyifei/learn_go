package main

import(
	UseMysql "use_mysql/practice"
)

func main() {
    /*
	dbAccess := UseMysql.ConnectToDB()
	dbAccess.InsertOne()
	dbAccess.SelectMultiUser()
	dbAccess.SelectOneUser(2001)
	dbAccess.UpdateRecord(2001)
	dbAccess.DeleteRecord(2001)
	dbAccess.ExecTrans()
    **/


    //测试数据库连接池
	//UseMysql.SelectMultiUser(1)
    UseMysql.TestConnectionPool(1000)

    defer UseMysql.SqlDB.Close()
}