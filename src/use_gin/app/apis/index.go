package apis

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	redis "github.com/go-redis/redis/v7"
	"net/http"
	"strconv"
	"use_gin/app/common"
	"use_gin/app/models"
	"use_gin/app/services"
)

func IndexApi(c *gin.Context) {
	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.56.102:6379",
		Password: "",
		DB:       0,
	})
	name, _ := client.Get("name").Result()
	c.String(http.StatusOK, name+"It works")
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
		c.JSON(http.StatusOK, gin.H{"msg": msg})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": msg})
	return
}

func CheckToken(c *gin.Context) {
	token := c.Request.FormValue("token")
	loginService := services.JwtLoginService{}
	err := loginService.GetUserInfoByToken(token)
	msg := "ok"
	if err != nil {
		msg = err.Error()
		c.JSON(http.StatusOK, gin.H{"msg": msg})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": msg})
	return
}

func AddUser(c *gin.Context) {
	userName := c.Request.FormValue("username")
	email := c.Request.FormValue("email")

	user := models.Users{UserName: userName, Email: email}
	ret, err := user.AddUser()
	common.CheckErr(err)
	msg := fmt.Sprintf("insert successful %d", ret)
	c.JSON(http.StatusOK, gin.H{"msg": msg})
}

func EditUser(c *gin.Context) {
	userName := c.Request.FormValue("username")
	email := c.Request.FormValue("email")
	id, err := strconv.ParseInt(c.Request.FormValue("id"), 10, 64)
	common.CheckErr(err)
	user := models.Users{UserName: userName, Email: email, Id: id}
	ret, err := user.Update()
	common.CheckErr(err)
	msg := fmt.Sprintf("affected rows:%d", ret)
	c.JSON(http.StatusOK, gin.H{"msg": msg})
}

func DelUser(c *gin.Context) {
	id := c.Param("ID")
	// id, err := strconv.ParseInt(c.Request.FormValue("id"), 10, 64)
	// common.CheckErr(err)
	Id, err := strconv.Atoi(id)
	common.CheckErr(err)
	user := &models.Users{Id: int64(Id)}

	//根据id进行查询，如果未查到则抛出错误信息
	user, err = user.GetUserById(user.Id)

	fmt.Println("user:", user)
	if err != nil {
		err = errors.New("invalid parameter Id:" + id)
		c.JSON(http.StatusOK, gin.H{"msg": err.Error()})
		return
	}

	ret, err := user.Delete()
	common.CheckErr(err)
	msg := fmt.Sprintf("delete successful, affected rows: %d", ret)
	c.JSON(http.StatusOK, gin.H{"msg": msg})
}

func ListUser(c *gin.Context) {
	req := c.Request
	fmt.Println(req)

	user := models.Users{}
	var userArr []models.Users
	userArr = make([]models.Users, 0)
	userArr, err := user.GetUsers()
	common.CheckErr(err)
	msg := ""
	c.JSON(http.StatusOK, gin.H{"msg": msg, "data": userArr})
}
