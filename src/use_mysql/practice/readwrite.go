package practice

import (
	"database/sql"
	"fmt"
	_ "mysql"
	"strconv"
)

type User struct {
	id int
	username string
	email string
}

func connectToDB() *sql.DB {
	db, err := sql.Open("mysql", "root:SYF!123mysql@tcp(127.0.0.1:3306)/test?charset=utf8")
	checkErr(err)
	return db
}

func InsertToDB() {
	db := connectToDB()
	stmt, err := db.Prepare("insert into users(username,email) values(?,?)")
	checkErr(err)
	res,err := stmt.Exec("caihewei", "caihewei@yonyou.com")
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println("inserted id is:", id)
	fmt.Println("insert operation is over")
	defer db.Close()
}

func SelectMultiUser() {
	db := connectToDB()
	rows, err := db.Query("select id,username from users order by id desc limit 0,?",5)
	checkErr(err)
	var id int
	var username string
	for rows.Next() {
		rerr := rows.Scan(&id, &username)
		checkErr(rerr)
		fmt.Printf("id:%s,username:%s \n", strconv.Itoa(id), username)
	}
	defer rows.Close()
	defer db.Close()
}

func UpdateRecord() {

}

func DeleteRecord() {

}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
