/**
 * init()函数的执行顺序，首先会先执行引用包里的init()函数，再执行main包里的init()的函数
 * 本示例演示了如果使用数据库连接池时，数据库里的在线客户端情况
 */
package main

import(
	"sync"
	UseMysql "use_mysql/practice"
	"fmt"
)

/**
 * 
 * @return {[type]} [description]
 */
func init() {
	fmt.Println("in use_mysql.go, init()")
}

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

    //UseMysql.Table2Struct()



    //测试数据库连接池
	//UseMysql.SelectMultiUser(1)

	var wg sync.WaitGroup
    wg.Add(4)
    go func() {
    	defer wg.Done()
    	UseMysql.TestConnectionPool(100)
	}()
	go func() {
		defer wg.Done()
		UseMysql.TestConnectionPool(200)
	}()
	go func() {
		defer wg.Done()
		UseMysql.TestConnectionPool(300)
	}()
	go func() {
		defer wg.Done()
		UseMysql.TestConnectionPool(400)
	}()
    wg.Wait()
	fmt.Println("cmplete")

    //注意这里的defer语句，程序运行结束后一定要关闭连接
    defer UseMysql.SqlDB.Close()
}