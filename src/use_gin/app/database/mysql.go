package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"use_gin/app/common"
	"use_gin/app/config"
)

var SqlDB * sql.DB

/**
	初始化数据库连接
 */
func init() {
	var err error
	SqlDB, err = sql.Open("mysql", fmt.Sprintf("root:SYF!123mysql@tcp(%s:3306)/test", config.ServerHost))
	common.CheckErr(err)
	err = SqlDB.Ping()
	common.CheckErr(err)

	SqlDB.SetMaxOpenConns(5)
	SqlDB.SetMaxIdleConns(1)

	//SqlDB.SetConnMaxLifetime(time.Second*10)
}

