package core

import (
	"blog_server/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

func InitGorm() {
	if global.CONFIG.Mysql.Host == "" {
		global.LOG.Warning("host为空")
	}
	dsn := global.CONFIG.Mysql.Dsn()
	var mysqlLogger logger.Interface
	if global.CONFIG.System.Env == "dev" {
		// 开发者环境显示所有sql
		mysqlLogger = logger.Default.LogMode(logger.Info)
	} else {
		mysqlLogger = logger.Default.LogMode(logger.Error)
	}
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: mysqlLogger})
	if err != nil {
		global.LOG.Panicf("打开数据库连接失败: %s", err)
	}
	sqlDb, _ := db.DB()
	// SetMaxIdleConns: 设置空闲连接池中链接的最大数量
	sqlDb.SetMaxIdleConns(global.CONFIG.Mysql.MaxIdleConns)
	// SetMaxOpenConns: 设置打开数据库链接的最大数量
	sqlDb.SetMaxOpenConns(global.CONFIG.Mysql.MaxOpenConns)
	// SetConnMaxLifetime: 设置链接可复用的最大时间
	sqlDb.SetConnMaxLifetime(5 * time.Minute)
	global.DB = db
	global.LOG.Info("gorm init success!")
}
