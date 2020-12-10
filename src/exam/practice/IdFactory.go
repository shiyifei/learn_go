/**
 * @Author:shiyf
 * @Date: 2020/12/8 22:48
 **/

package practice

import (
	"fmt"
	"time"
)

type IdFactory struct {
	workerId int64
	dataCenterId int64
	sequence int64
	beginTime int64
}

func init() {
	fmt.Println(time.Now().Unix())
	fmt.Println(time.Now().UnixNano())
	//毫秒级的整数
	fmt.Println(time.Now().UnixNano()/1000000)

	//将字符串转为时间戳
	var strTime string = "2020-12-08 00:00:00"
	var standard string = "2006-01-02 15:04:05"
	time1,_ := time.Parse(standard, strTime)
	fmt.Println(time1.UTC())
}

func GenerateId() {

}