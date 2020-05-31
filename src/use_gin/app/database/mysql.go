package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"use_gin/app/common"
)

var SqlDB * sql.DB

/**
	初始化数据库连接
 */
func init() {
	var err error
	SqlDB, err = sql.Open("mysql", "root:SYF!123mysql@tcp(192.168.1.102:3306)/test")
	common.CheckErr(err)
	err = SqlDB.Ping()
	common.CheckErr(err)

	SqlDB.SetMaxOpenConns(5)
	SqlDB.SetMaxIdleConns(1)

	//SqlDB.SetConnMaxLifetime(time.Second*10)
}

