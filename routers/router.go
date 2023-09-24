package routers

import (
	v1 "blog_server/api/v1"
	"blog_server/global"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode(global.CONFIG.System.Env)
	r := gin.Default()
	router := r.Group("api/v1")
	// 用户模块的接口
	user := router.Group("user")
	{
		user.POST("add", v1.AddUser)
		user.GET("info", v1.GetUserInfo)
		user.GET("exist", v1.UserIfExist)
		user.GET("list", v1.GetUserList)
		user.DELETE(":id", v1.DeleteUser)
		user.PUT(":id", v1.UpdateUser)
	}
	// 分类模块的接口
	cate := router.Group("category")
	{
		cate.POST("add", v1.AddCategory)
		cate.GET("info", v1.GetCategoryInfo)
		cate.GET("list", v1.GetCategoryList)
		cate.DELETE(":id", v1.DeleteCategory)
		cate.PUT(":id", v1.UpdateCategory)
	}
	return r
}
