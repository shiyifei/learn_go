package practice

import (
	"fmt"
	"math/rand"
	"database/sql"
    _ "github.com/go-sql-driver/mysql"

	"strconv"
)

type DBAccess struct {
	Db *sql.DB
}

type User struct {
	id int
	username string
	email string
	createtime string
}


func ConnectToDB() *DBAccess{
	db, err := sql.Open("mysql", "root:SYF!123mysql@tcp(127.0.0.1:3306)/test?charset=utf8")
	checkErr(err)
	return &DBAccess{db}
}


/*
	插入单条记录
 */
func (worker *DBAccess) InsertOne() {
	stmt, err := worker.Db.Prepare("insert into users(username,email) values(?, ?)")
	checkErr(err)
	res,err := stmt.Exec("caihewei", "caihewei@yonyou.com")
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println("last inserted id is:", id)

	rowsAffected, err := res.RowsAffected()
	checkErr(err)
	fmt.Println("Rows Affected:", rowsAffected)

	fmt.Println("insert operation is over")
	//defer worker.Db.Close()
}

/* 查询多条记录 */
func (worker *DBAccess) SelectMultiUser() {
	rows, err := worker.Db.Query("select * from users order by id desc limit 0,?",5)
	checkErr(err)

	columns, err := rows.Columns()
	checkErr(err)
	scanArgs := make([]interface{}, len(columns))
	values := make([][]byte, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}

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
	查询单条记录
 */
func (worker *DBAccess) SelectOneUser(id int) {
	row := worker.Db.QueryRow("select id,username,email,createtime from users where id = ?",id)
	user := User{}
	err := row.Scan(&user.id, &user.username, &user.email, &user.createtime)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(user)
	}


}

/**
	更新记录
 */
func (worker *DBAccess) UpdateRecord(id int) {
	stmt, err := worker.Db.Prepare("update users set username=? where id = ?")
	checkErr(err)
	ret, err := stmt.Exec("caihewei" + strconv.Itoa(rand.Intn(1000)),id)
	checkErr(err)
	rowsAffected, err := ret.RowsAffected()
	checkErr(err)
	fmt.Println("after update Rows Affected:", rowsAffected)
	worker.SelectOneUser(id)
}

/**
	删除记录
 */
func (worker *DBAccess) DeleteRecord(id int) {
	stmt, err := worker.Db.Prepare("delete from users where id = ?")
	checkErr(err)
	ret, err := stmt.Exec(id)
	checkErr(err)
	rowsAffected, err := ret.RowsAffected()
	checkErr(err)
	fmt.Println("after deleting,Rows Affected:", rowsAffected)
	worker.SelectOneUser(id)
}

/*
	测试数据库事务
 */
func (worker *DBAccess) ExecTrans() {
	tx, err := worker.Db.Begin()
	checkErr(err)

	ret1, err := tx.Exec("update users set username=\"guoxiaofeng\" where id=?",1000)
	if err != nil {
		panic(err)
		defer tx.Rollback()
	}
	num1, err := ret1.RowsAffected()
	if err != nil {
		panic(err)
		defer tx.Rollback()
	}
	ret2, err := tx.Exec("update users set username=\"guoke\" where id=?",1001)
	if err != nil {
		panic(err)
		defer tx.Rollback()
	}
	num2, err := ret2.RowsAffected()
	if err != nil {
		panic(err)
		defer tx.Rollback()
	}
	if num1>0 && num2>0 {
		tx.Commit()
	} else {
		tx.Rollback()
	}
}

/**
	程序中断抛出异常
 */
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
