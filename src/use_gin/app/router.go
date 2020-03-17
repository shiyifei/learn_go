package app

import (
	"github.com/gin-gonic/gin"
	. "use_gin/app/apis"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", IndexApi)
	router.GET("/listUsers", ListUser)
	router.POST("/addUser", AddUser)
	router.PUT("/editUser", EditUser)
	router.DELETE("/delUser", DelUser)
	router.POST("/delUser", DelUser)
	return router
}