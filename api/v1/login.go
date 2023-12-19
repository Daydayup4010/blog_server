package v1

import (
	"blog_server/middleware"
	"blog_server/models"
	"blog_server/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	var user models.User
	var code int
	var token string
	c.ShouldBindJSON(&user)
	code = models.VerifyLogin(user.Username, user.Password)
	if code == errmsg.SUCCESS {
		token, code = middleware.CreateToken(user.Username)
	}
	c.JSON(http.StatusOK, gin.H{
		"code":  code,
		"msg":   errmsg.GetErrMsg(code),
		"token": token,
	})
}
