package v1

import (
	"blog_server/models"
	"blog_server/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// AddArt 添加文章
func AddArt(c *gin.Context) {
	var art models.Article
	_ = c.ShouldBindJSON(&art)
	code := models.CreateArt(&art)
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"data":    art,
		"message": errmsg.GetErrMsg(code),
	})

}

// 查询文章详情信息

// GetArtList 查询文章列表
func GetArtList(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	list, code := models.GetArtList(pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"data":    list,
		"total":   len(list),
		"message": errmsg.GetErrMsg(code),
	})
}

// UpdateArt 编辑文章
func UpdateArt(c *gin.Context) {
	var data models.Article
	c.ShouldBindJSON(&data)
	id, _ := strconv.Atoi(c.Param("id"))
	code := models.UpdateArt(id, &data)
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": errmsg.GetErrMsg(code),
	})

}

// DeleteArt 删除文章
func DeleteArt(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := models.DeleteArt(id)
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": errmsg.GetErrMsg(code),
	})
}
