package main

import (
	"blog_server/core"
	"blog_server/global"
	"blog_server/routers"
)

func main() {
	// 读取初始化文件
	core.InitConf()
	core.InitLogger()
	core.InitGorm()
	engine := routers.InitRouter()
	addr := global.CONFIG.Server.GetAddr()
	err := engine.Run(addr)
	if err != nil {
		global.LOG.Panicf("服务启动失败: %s", err)
	}
}
