package routers

import (
	v1 "blog_server/api/v1"
	"blog_server/global"
	"blog_server/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode(global.CONFIG.Server.Env)
	r := gin.New()
	r.Use(gin.Recovery(), middleware.Logger())
	authn := r.Group("api/v1")

	authn.POST("upload", v1.Upload)
	public := r.Group("api/v1")
	public.Use(middleware.JwtToken())
	// 用户模块的接口
	user := authn.Group("user")
	{
		user.POST("add", v1.AddUser)
		user.DELETE(":id", v1.DeleteUser)
		user.PUT(":id", v1.UpdateUser)
	}
	user2 := public.Group("user")
	{
		user2.GET("info", v1.GetUserInfo)
		user2.GET("exist", v1.UserIfExist)
		user2.GET("list", v1.GetUserList)
		user2.POST("login", v1.Login)
	}

	// 分类模块的接口
	cate := authn.Group("category")
	{
		cate.POST("add", v1.AddCategory)
		cate.DELETE(":id", v1.DeleteCategory)
		cate.PUT(":id", v1.UpdateCategory)
	}
	cate2 := public.Group("category")
	{
		cate2.GET("list", v1.GetCategoryList)
	}

	// 文章模块的接口
	art := authn.Group("article")
	{
		art.POST("add", v1.AddArt)
		art.DELETE(":id", v1.DeleteArt)
		art.PUT(":id", v1.UpdateArt)

	}
	art2 := public.Group("article")
	{
		art2.GET("list", v1.GetArtList)
		art2.GET("info", v1.GetArtInfo)
		art2.GET("category/:id/list", v1.GetCateArtList)
	}
	return r
}
