package v1

import (
	"blog_server/models"
	"blog_server/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// UserIfExist 查询用户是否存在
func UserIfExist(c *gin.Context) {
}

// AddUser 添加用户
func AddUser(c *gin.Context) {
	var user models.User
	_ = c.ShouldBindJSON(&user)
	code := models.CheckUser(user.Username)
	if code == errmsg.SUCCESS {
		models.CreateUser(&user)
		c.JSON(http.StatusOK, gin.H{
			"code":    code,
			"data":    user,
			"message": errmsg.GetErrMsg(code),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    code,
			"message": errmsg.GetErrMsg(code),
		})
	}

}

// GetUserInfo 查询用户详情信息
func GetUserInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	info, code := models.GetUserInfo(id)
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"data":    info,
		"message": errmsg.GetErrMsg(code),
	})
}

// GetUserList 查询用户列表
func GetUserList(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))

	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	users, code := models.GetUsers(pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"data":    users,
		"message": errmsg.GetErrMsg(code),
	})
}

// UpdateUser 编辑用户
func UpdateUser(c *gin.Context) {
	var user models.User
	id, _ := strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&user)
	code := models.CheckUserFromId(id)
	code = models.CheckUser(user.Username)
	if code == errmsg.SUCCESS {
		code = models.UpdateUser(id, &user)
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": errmsg.GetErrMsg(code),
	})

}

// DeleteUser 删除用户
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := models.DeleteUser(id)
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": errmsg.GetErrMsg(code),
	})
}
