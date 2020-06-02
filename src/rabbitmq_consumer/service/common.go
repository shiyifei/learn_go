package service

import(
"log"
)



/**
	输出异常信息
 */
func FailOnError(err error, msg string) {
	if err != nil {
		log.Printf("%s: %s \n", msg, err)
	}
}
