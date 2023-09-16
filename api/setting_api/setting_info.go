package setting_api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (SettingApi) GetSettingInfo(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"msg": "success"})
}
