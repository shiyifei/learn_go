package apis

/**
 * 演示如何操作cookie,注意这里的域名要跟访问地址一致。本机IP：192.168.56.107
 */

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ReadCookie(c *gin.Context) {
	val, _ := c.Cookie("name")
	c.String(http.StatusOK, "Cookie:%s", val)
}

func WriteCookie(c *gin.Context) {
	c.SetCookie("name", "shiyifei", 10, "/", "", false, true)
}

func ClearCookie(c *gin.Context) {
	c.SetCookie("name", "shiyifei", -1, "/", "192.168.56.107", false, true)
}
