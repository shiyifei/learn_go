package app

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	. "use_gin/app/apis"

	"fmt"
	"time"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	// LoggerWithFormatter middleware will write the logs to gin.DefaultWriter
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	//简单的路由组 v1
	v1 := router.Group("/user")
	{
		v1.GET("/", IndexApi)
		v1.GET("/listUsers", ListUser)
		v1.POST("/addUser", AddUser)
		v1.PUT("/editUser", EditUser)
		v1.DELETE("/delUser/:ID", DelUser)
		v1.POST("/delUser/:ID", DelUser)
	}

	v2 := router.Group("/home")
	{
		v2.POST("/user/login", UserLogin)
		v2.POST("/user/check", CheckToken)
	}

	v3 := router.Group("/cookie")
	{
		v3.GET("/write", WriteCookie)
		v3.GET("/read", ReadCookie)
		v3.GET("/clear", ClearCookie)
	}

	//defer common.TryRecover()

	return router
}
