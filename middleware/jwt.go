package middleware

import (
	"blog_server/global"
	"blog_server/utils/errmsg"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

// var JwtKey = []byte(global.CONFIG.Server.JwtKey)
var JwtKey = []byte("dad131d1")

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// CreateToken 生成token
func CreateToken(username string) (string, int) {
	expireTime := time.Now().Add(time.Hour * 10)
	SetClaims := MyClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "ginBlog",
		},
	}
	reqClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, SetClaims)
	token, err := reqClaim.SignedString(JwtKey)
	if err != nil {
		global.LOG.Info(err)
		return "", errmsg.ERROR
	}
	return token, errmsg.SUCCESS
}

// VerifyToken 验证token
func VerifyToken(token string) (*MyClaims, int) {
	setToken, _ := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if key, ok := setToken.Claims.(*MyClaims); ok && setToken.Valid {
		return key, errmsg.SUCCESS
	} else {
		return nil, errmsg.ERROR
	}

}

// JwtToken jwt中间件
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHerder := c.Request.Header.Get("Authorization")
		code := errmsg.SUCCESS
		if tokenHerder == "" {
			code = errmsg.ERROR_TOKEN_EXIST
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		verifyToken := strings.SplitN(tokenHerder, " ", 2)
		if len(verifyToken) != 2 && verifyToken[0] != "Bearer" {
			code = errmsg.ERROR_TOKEN_TYPE_WRONG
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		key, tCode := VerifyToken(verifyToken[1])
		if tCode == errmsg.ERROR {
			code = errmsg.ERROR_TOKEN_WRONG
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		if time.Now().Unix() > key.ExpiresAt {
			code = errmsg.ERROR_TOKEN_TIMEOUT
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		c.Set("username", key.Username)
		c.Next()
	}
}
