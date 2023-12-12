package v1

import (
	"blog_server/models"
	"blog_server/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// AddCategory 添加分类
func AddCategory(c *gin.Context) {
	var cate models.Category
	_ = c.ShouldBindJSON(&cate)
	code := models.CheckCategory(cate.Name)
	if code == errmsg.SUCCESS {
		models.CreateCategory(&cate)
		c.JSON(http.StatusOK, gin.H{
			"code":    code,
			"data":    cate,
			"message": errmsg.GetErrMsg(code),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    code,
			"message": errmsg.GetErrMsg(code),
		})
	}

}

// GetCategoryInfo 查询分类详情信息
//func GetCategoryInfo(c *gin.Context) {
//	id, _ := strconv.Atoi(c.Query("id"))
//	info, code := models.GetUserInfo(id)
//	c.JSON(http.StatusOK, gin.H{
//		"code":    code,
//		"data":    info,
//		"message": errmsg.GetErrMsg(code),
//	})
//}

// GetCategoryList 获取分类列表
func GetCategoryList(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))

	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	users, code := models.GetCategoryList(pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"data":    users,
		"message": errmsg.GetErrMsg(code),
	})
}

// UpdateCategory 更新分类
func UpdateCategory(c *gin.Context) {
	var cate models.Category
	id, _ := strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&cate)
	code := models.CheckCategoryFromId(id)
	code = models.CheckCategory(cate.Name)
	if code == errmsg.SUCCESS {
		code = models.UpdateCategory(id, &cate)
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": errmsg.GetErrMsg(code),
	})

}

// DeleteCategory 删除分类
func DeleteCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := models.DeleteCategory(id)
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": errmsg.GetErrMsg(code),
	})
}
