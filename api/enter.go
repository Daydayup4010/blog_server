package api

import "blog_server/api/setting_api"

type ApiGroup struct {
	SettingApi setting_api.SettingApi
}

var ApiGroupApp = new(ApiGroup)
