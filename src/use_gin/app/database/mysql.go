package database

import (
	"database/sql"
	"use_gin/app/common"
	_ "github.com/go-sql-driver/mysql"
)

var SqlDB * sql.DB

/**
	初始化数据库连接
 */
func init() {
	var err error
	SqlDB, err = sql.Open("mysql", "manager:SYF!123mysql@tcp(192.168.56.102:3306)/test")
	common.CheckErr(err)
	err = SqlDB.Ping()
	common.CheckErr(err)
}

