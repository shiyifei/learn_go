package models

import (
	"errors"
	"fmt"
	"use_gin/app/common"
	db "use_gin/app/database"
)

type Users struct {
	Id         int64  `json:"id" form:"id"`
	UserName   string `json:"username" form:"username"`
	Email      string `json:"email" form:"email"`
	CreateTime string `json:"createtime"`
}

func (p *Users) CheckUserPwd(username, password string) (Users, error) {
	rows, err := db.SqlDB.Query("select id,username,email,createtime from users where username=? and userpwd=?", username, password)

	fmt.Println("in CheckUserPwd ", rows, err, " end ")
	defer rows.Close()
	common.CheckErr(err)
	var user Users

	var count int = 0
	for rows.Next() {
		count += 1
		err = rows.Scan(&user.Id, &user.UserName, &user.Email, &user.CreateTime)
		common.CheckErr(err)
	}
	if count == 0 {
		return user, errors.New("no data")
	}
	err = rows.Err()
	fmt.Println("in CheckUserPwd  rows.Err:", err, " end ")

	return user, nil
}

func (p *Users) GetUserById(id int64) (*Users, error) {
	rows, err := db.SqlDB.Query("select id,username,email,createtime from users where id=?", id)

	fmt.Println("in GetUserById ", rows, err, " end ")
	defer rows.Close()
	common.CheckErr(err)
	var user Users

	var count int = 0
	for rows.Next() {
		count += 1
		err = rows.Scan(&user.Id, &user.UserName, &user.Email, &user.CreateTime)
		common.CheckErr(err)
	}
	if count == 0 {
		return &user, errors.New("no data")
	}

	return &user, nil
}

func (p *Users) GetUserByName(username string) (Users, error) {
	rows, err := db.SqlDB.Query("select id,username,email,createtime from users where username=? limit 1", username)
	defer rows.Close()
	common.CheckErr(err)
	var user Users
	for rows.Next() {
		err = rows.Scan(&user.Id, &user.UserName, &user.Email, &user.CreateTime)
		common.CheckErr(err)
	}
	return user, nil
}

func (p *Users) GetUsers() ([]Users, error) {
	rows, err := db.SqlDB.Query("select id,username,email,createtime from users order by id desc limit 0,10")
	defer rows.Close()
	common.CheckErr(err)

	var records []Users
	records = make([]Users, 0)

	//time.Sleep(time.Second *5)

	for rows.Next() {
		var user Users
		err = rows.Scan(&user.Id, &user.UserName, &user.Email, &user.CreateTime)
		common.CheckErr(err)
		records = append(records, user)
	}
	return records, nil
}

func (p *Users) AddUser() (int64, error) {
	stmt, err := db.SqlDB.Prepare("insert into users(username, email) values(?, ?)")
	common.CheckErr(err)
	res, err := stmt.Exec(p.UserName, p.Email)
	common.CheckErr(err)
	num, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return num, err
}

func (p *Users) Update() (int64, error) {
	stmt, err := db.SqlDB.Prepare("update users set email=?,username=? where id=?")
	common.CheckErr(err)
	res, err := stmt.Exec(p.Email, p.UserName, p.Id)
	common.CheckErr(err)

	num, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return num, err
}

func (p *Users) Delete() (int64, error) {
	stmt, err := db.SqlDB.Prepare("delete from users where id=?")
	common.CheckErr(err)
	res, err := stmt.Exec(p.Id)
	common.CheckErr(err)
	num, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return num, err
}
