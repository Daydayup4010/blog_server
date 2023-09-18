package setting_api

import (
	"blog_server/models/res"
	"github.com/gin-gonic/gin"
)

func (SettingApi) GetSettingInfo(ctx *gin.Context) {
	res.Ok(map[string]any{"id": 131}, "success", ctx)
}
