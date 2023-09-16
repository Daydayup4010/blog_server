package global

import (
	"blog_server/config"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	CONFIG *config.Config // 配置文件
	DB     *gorm.DB       // 数据库连接对象
	LOG    *logrus.Logger
)
