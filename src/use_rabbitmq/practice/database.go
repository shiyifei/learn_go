package practice

import(
"database/sql"
_ "github.com/go-sql-driver/mysql"
"fmt"
)

var SqlDB *sql.DB

func init() {
	fmt.Println("in access.go, init()")
	var err error
	localIp := GetLocalIp()
	/*if localIp == "192.168.56.107" {
		localIp = "192.168.56.110"
	} else {
		localIp = "192.168.1.102"
	}*/

	localIp = "192.168.56.105"

	SqlDB, err = sql.Open("mysql", fmt.Sprintf("root:SYF!123mysql@tcp(%s:3306)/test", localIp))
	FailOnError(err, "can not connect database")
	err = SqlDB.Ping()
	FailOnError(err, "ping sqldb error")

	SqlDB.SetMaxOpenConns(4)
	SqlDB.SetMaxIdleConns(2)

	//SqlDB.SetConnMaxLifetime(time.Second*10)
}