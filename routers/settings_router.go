package routers

import (
	"blog_server/api"
	"github.com/gin-gonic/gin"
)

func SettingRouter(group *gin.RouterGroup) {
	settingApi := api.ApiGroupApp.SettingApi
	settingGroup := group.Group("/setting")
	{
		settingGroup.GET("/info", settingApi.GetSettingInfo)
	}
}
