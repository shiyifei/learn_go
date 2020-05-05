package practice

import(
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
    "time"
	"fmt"
	"net"
	"strings"
)

var SqlDB *sql.DB

func init() {
	fmt.Println("in access.go, init()")
	var err error
	localIp := GetLocalIp()
	if localIp == "192.168.56.107" {
		localIp = "192.168.56.102"
	}

	SqlDB, err = sql.Open("mysql", fmt.Sprintf("root:SYF!123mysql@tcp(%s:3306)/test", localIp))
	checkErr(err)
	err = SqlDB.Ping()
	checkErr(err)

    SqlDB.SetMaxOpenConns(4)
    SqlDB.SetMaxIdleConns(2)

    //SqlDB.SetConnMaxLifetime(time.Second*10)
}

func GetLocalIp() string {
	conn, err := net.Dial("udp", "www.baidu.com:80")
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().String()

	localIp := strings.Split(localAddr, ":")[0]
	fmt.Println("localAddr:", localAddr, "localIP:",localIp)
	return localIp
}


//测试读数据和连接池使用效果
func ReadData(offset int) []map[string]string {
	rows, err := SqlDB.Query("select * from users order by id desc limit ?,10", offset)
	checkErr(err)
	columns, _ := rows.Columns()

    fmt.Println(columns)
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}
    //time.Sleep(time.Second * 5) //测试连接池效果，保持db连接不释放

	var records []map[string]string
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		record := make(map[string]string)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}
		records = append(records, record)
	}
	return records
}

func SelectMultiUser(offset int) {
	rows, err := SqlDB.Query("select * from users order by id desc limit ?,10",offset)
	checkErr(err)

	columns, err := rows.Columns()
	checkErr(err)
	scanArgs := make([]interface{}, len(columns))
	values := make([][]byte, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	time.Sleep(time.Second * 5) //测试连接池效果，保持db连接不释放
	for rows.Next() {
		rerr := rows.Scan(scanArgs...)
		checkErr(rerr)
		//单行
		record := make(map[string]string)
		for k,v := range values {
			if v != nil {
				record[columns[k]] = string(v)
			}
		}
		fmt.Println(record)
	}
	defer rows.Close()
	//defer worker.Db.Close()
}

/*
    测试是否使用连接池
 */
func TestConnectionPool(page int) {
    for page=0; page<10; page++ {
		SelectMultiUser(page)
    }
}

func Insert(username, email string) (int64, error) {
	stmt, err := SqlDB.Prepare("insert into users(username, email) values(?, ?)")
	checkErr(err)
	res, err := stmt.Exec(username, email)
	checkErr(err)
	num, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return num, err
}

func Update(id int) (int64, error) {
	stmt, err := SqlDB.Prepare("update users set email=? where id=?")
	checkErr(err)
	res, err := stmt.Exec( "wuzhenghao@yxp.com", id)
	checkErr(err)

	num, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return num, err
}

func Delete(id int) (int64, error) {
	stmt, err := SqlDB.Prepare("delete from users where id=?")
	checkErr(err)
	res, err := stmt.Exec(id)
	checkErr(err)
	num, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return num, err
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}