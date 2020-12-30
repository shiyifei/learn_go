package practice

import(
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"time"
	"fmt"
)

var SqlDB *sql.DB

func init() {
	var err error
	SqlDB, err = sql.Open("mysql", fmt.Sprintf("root:SYF!123mysql@tcp(%s:3306)/test", "127.0.0.1"))
	CheckErr(err)
	err = SqlDB.Ping()
	CheckErr(err)

	SqlDB.SetMaxOpenConns(4)
	SqlDB.SetMaxIdleConns(2)
	//SqlDB.SetConnMaxLifetime(time.Second*10)
}

/**
 * @Description: 处理预订单数据，更新预订单表、插入订单表
 * @param data
 */
func ProcessData(data map[string]interface{}) {
	tmp, ok := data["pre_order_id"].(float64)
	preOrderId := int(tmp)

	fmt.Printf("传入的pre_order_id:%T, 转化后：%#v, ok:%#v \n", data["pre_order_id"], preOrderId, ok)
	//判断是否有该预订单id
	rows, err := SqlDB.Query("select count(*) num from biz_pre_order where id=? and status=1", preOrderId)
	CheckErr(err)
	retArr := selectRows(rows)
	if len(retArr) == 0 {
		error := errors.New(fmt.Sprintf("传入的数据非法, pre_order_id:%v \n", preOrderId))
		CheckErr(error)
	} else {
		num, err := strconv.Atoi(retArr[0]["num"])
		if err != nil || num <=0 {
			err = errors.New(fmt.Sprintf("未查到该数据, pre_order_id:%d \n", preOrderId))
			CheckErr(err)
		}
	}

	//数据库开始事务
	tx, err := SqlDB.Begin()
	if err != nil {
		if tx != nil {
			_ = tx.Rollback()
		}
		error := errors.New(fmt.Sprintf("begin trans action failed, err:%v\n", err))
		CheckErr(error)
		return
	}

	//插入订单表
	ret, err := Insert(tx, data)
	if err != nil {
		tx.Rollback()
		return
	}
	if ret <= 0 {
		tx.Rollback()
		return
	}

	//修改预订单表
	ret, err = Update(tx, preOrderId)
	if err != nil {
		tx.Rollback()
		return
	}
	if ret == 0 {
		tx.Rollback()
		return
	}

	//数据库结束事务
	tx.Commit()
	return
}

/**
 * @Description: 查询数据库
 * @param rows 查询后的记录 类型 *sql.Rows
 * @return []map[string]string
 */
func selectRows(rows *sql.Rows) []map[string]string {
	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([][]byte, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	var records []map[string]string
	var err error
	record := make(map[string]string)
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		CheckErr(err)
		for k, v := range values {
			if v != nil {
				record[columns[k]] = string(v)
			}
		}
		records = append(records, record)
	}
	return records
}

/**
 * @Description: 读取多条数据
 * @param limit
 * @return []map[string]string
 */
func ReadData(limit int) []map[string]string {
	rows, err := SqlDB.Query("select * from biz_pre_order order by id desc limit 0,?", limit)
	CheckErr(err)
	columns, _ := rows.Columns()

	//fmt.Println(columns)
	scanArgs := make([]interface{}, len(columns))
	values := make([][]byte, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	var records []map[string]string
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		record := make(map[string]string)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col)
			}
		}
		records = append(records, record)
	}
	defer rows.Close()
	return records
}

func Insert(tx * sql.Tx, data map[string]interface{}) (int64, error) {
	sqlInsert := "insert into biz_order(user_id,op_type,app_id,pre_order_id,created_at) values(?,?,?,?,?)"
	fUserId,_ := data["user_id"].(float64)
	userId := int(fUserId)
	fOpType,_ := data["op_type"].(float64)
	opType := int(fOpType)
	fAppId,_ := data["app_id"].(float64)
	appId := int(fAppId)
	fPreOrderId,_ := data["pre_order_id"].(float64)
	preOrderId := int(fPreOrderId)
	createdAt := time.Now().Unix()

	res, err := tx.Exec(sqlInsert, userId, opType, appId, preOrderId, createdAt)
	CheckErr(err)
	num, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return num, err
}

func Update(tx *sql.Tx, id int) (int64, error) {
	sql := "update biz_pre_order set status=4, updated_at=? where id=?"
	res, err := tx.Exec(sql, time.Now().Unix(), id)
	CheckErr(err)
	num, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return num, err
}

/**
 * @Description: 根据id删除行
 * @param int id id
 * @return int64
 * @return error
 */
func Delete(id int) (int64, error) {
	stmt, err := SqlDB.Prepare("delete from biz_pre_order where id=?")
	CheckErr(err)
	res, err := stmt.Exec(id)
	CheckErr(err)
	num, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return num, err
}

/**
 * @Description: 检查错误
 * @param err
 */
func CheckErr(err error) {
	if err != nil {
		defer MyRecover()
		panic(err)
	}
}

/**
 * @Description:  panic恐慌捕获
 */
func MyRecover() {
	r := recover()
	if  err, ok := r.(error); ok {
		// 捕获异常
		fmt.Println("捕获到错误：", err)
	} else {
		// 重新抛出一个异常
		//panic(r)
	}
}