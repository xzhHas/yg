package global

import (
	"gorm.io/gorm"
	"xback/user_srv/config"
)

var (
	DB *gorm.DB

	ServerConfig config.ServerConfig

	NacosConfig config.NacosConfig
)
