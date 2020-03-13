package common

import "log"

/**
	抛出错误
 */
func CheckErr(err error) {
	if err != nil {
		log.Fatal(err.Error())
		panic(err)
	}
}

