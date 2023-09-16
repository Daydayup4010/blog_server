package routers

import (
	"blog_server/api"
	"github.com/gin-gonic/gin"
)

func SettingRouter(engine *gin.Engine) {
	settingApi := api.ApiGroupApp.SettingApi
	settingGroup := engine.Group("api/setting")
	{
		settingGroup.GET("/info", settingApi.GetSettingInfo)
	}
}
