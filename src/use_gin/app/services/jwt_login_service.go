package services

import(
	"bytes"
	"crypto"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
	"use_gin/app/config"
	"use_gin/app/models"
)

type JwtLoginService struct {

}

type LoginInfo struct {
	Username string `json:"username"`
	Email string `json:"email"`
	Authcode string	`json:"auth_code"`
}

type MyCustomClaims struct {
	LoginInfo
	jwt.StandardClaims
}

/**
	根据username、验证码来创建token
 */
func (p *JwtLoginService) CreateSign(username,email,authcode string) (string, error) {
	secret := config.JwtShaSecret
	loginInfo := LoginInfo{Username:username, Email:email, Authcode:authcode}
	claims := MyCustomClaims {
		loginInfo,
		jwt.StandardClaims {
			IssuedAt:time.Now().Unix(),
			ExpiresAt:time.Now().Add(24*time.Hour).Unix(), //24小时后过期
		},
	}

	tokenClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaim.SignedString(secret)
	return token, err
}

/**
	根据传入的token解析出登录的用户名
 */
func (p *JwtLoginService) parseSign(strToken string) (string, error) {
	token, err := jwt.ParseWithClaims(strToken, &MyCustomClaims{}, func(token *jwt.Token)(interface{}, error){
		return config.JwtShaSecret, nil
	})

	claims, ok := token.Claims.(*MyCustomClaims)
	if ok && token.Valid {
		fmt.Printf("username:%v, authcode:%v,expired_at:%v", claims.Username, claims.Authcode, claims.StandardClaims.ExpiresAt)
		return claims.Username, nil
	} else {
		fmt.Println(err)
		return "",err
	}
}

func sha1Encrypt(input string) string {
	buf := bytes.NewBufferString(input)
	h := crypto.SHA1.New()
	h.Write([]byte(buf.String()))
	return hex.EncodeToString(h.Sum(nil))
}

/**
	根据用户名来查找用户的相关信息
 */
func (p *JwtLoginService) UserLogin(username, password, authcode string) error {
	 password = sha1Encrypt(password)
	 fmt.Printf("password:%s \n", password)
	 var user = models.Users{}
	 user, err := user.CheckUserPwd(username, password)

	 fmt.Println(user, err)

	 if err != nil {
	 	return  errors.New("用户名或密码错误")
	 }

	 token, err := p.CreateSign(user.UserName, user.Email, authcode)
	 fmt.Println(token, err)
	 return nil
}


/**
	根据用户名来查找用户的相关信息
 */
func (p *JwtLoginService) GetUserInfoByToken(token string) error {
	user := models.Users{}
 	username,err := p.parseSign(token)
 	if err != nil {
 		return err
	}
 	user, err = user.GetUserByName(username)
 	fmt.Println(user, err)
 	return err
}
