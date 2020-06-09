package service


import (
"database/sql"
_ "github.com/go-sql-driver/mysql"
)

var SqlDB * sql.DB

/**
	初始化数据库连接
 */
func init() {
	var err error
	SqlDB, err = sql.Open("mysql", "root:SYF!123mysql@tcp(192.168.56.102:3306)/test")
	FailOnError(err, "can not connect to mysql server")
	err = SqlDB.Ping()
	FailOnError(err, "can not ping to mysql")

	SqlDB.SetMaxOpenConns(5)
	SqlDB.SetMaxIdleConns(1)

	//SqlDB.SetConnMaxLifetime(time.Second*10)
}
