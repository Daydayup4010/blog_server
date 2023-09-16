package routers

import (
	"blog_server/global"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode(global.CONFIG.System.Env)
	engine := gin.Default()
	SettingRouter(engine)
	return engine
}
