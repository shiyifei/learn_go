package practice

import(
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var SqlDB *sql.DB

func init() {
	var err error
	SqlDB, err = sql.Open("mysql", "manager:SYF!123mysql@tcp(192.168.56.102:3306)/test")
	checkErr(err)
	err = SqlDB.Ping()
	checkErr(err)
}

func ReadData() []map[string]string {
	rows, err := SqlDB.Query("select * from users order by id desc limit 0,10")
	checkErr(err)
	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}
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