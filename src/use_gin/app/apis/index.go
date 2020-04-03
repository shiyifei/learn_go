package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"use_gin/app/common"
	"use_gin/app/models"
	"use_gin/app/services"
)


func IndexApi(c *gin.Context) {
	c.String(http.StatusOK, "It works")
}

func UserLogin(c *gin.Context) {
	userName := c.Request.FormValue("username")
	password := c.Request.FormValue("password")
	authcode := c.Request.FormValue("checkcode")
	loginService := services.JwtLoginService{}
	err := loginService.UserLogin(userName, password, authcode)
	msg := "ok"
	if err != nil {
		msg = err.Error()
		c.JSON(http.StatusOK, gin.H{"msg":msg})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg":msg})
	return
}

func CheckToken(c *gin.Context) {
	token := c.Request.FormValue("token")
	loginService := services.JwtLoginService{}
	err := loginService.GetUserInfoByToken(token)
	msg := "ok"
	if err != nil {
		msg = err.Error()
		c.JSON(http.StatusOK, gin.H{"msg":msg})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg":msg})
	return
}

func AddUser(c *gin.Context) {
	userName := c.Request.FormValue("username")
	email := c.Request.FormValue("email")

	user := models.Users{UserName:userName, Email:email}
	ret, err := user.AddUser()
	common.CheckErr(err)
	msg := fmt.Sprintf("insert successful %d", ret)
	c.JSON(http.StatusOK, gin.H{"msg":msg})
}

func EditUser(c *gin.Context) {
	userName := c.Request.FormValue("username")
	email := c.Request.FormValue("email")
	id, err := strconv.ParseInt(c.Request.FormValue("id"), 10, 64)
	common.CheckErr(err)
	user := models.Users{UserName:userName, Email:email, Id:id}
	ret, err := user.Update()
	common.CheckErr(err)
	msg := fmt.Sprintf("affected rows:%d", ret)
	c.JSON(http.StatusOK, gin.H{"msg":msg})
}

func DelUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Request.FormValue("id"), 10, 64)
	common.CheckErr(err)
	user := models.Users{Id:id}
	ret, err := user.Delete()
	common.CheckErr(err)
	msg := fmt.Sprintf("delete successful, affected rows: %d", ret)
	c.JSON(http.StatusOK, gin.H{"msg":msg})
}

func ListUser(c *gin.Context) {
	user := models.Users{}
	var userArr []models.Users
	userArr = make([]models.Users, 0)
	userArr, err := user.GetUsers()
	common.CheckErr(err)
	msg := ""
	c.JSON(http.StatusOK, gin.H{"msg":msg, "data":userArr})
}

