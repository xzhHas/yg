package global

import (
	"gorm.io/gorm"
	"xback/userop_srv/config"
)

var (
	DB *gorm.DB

	ServerConfig config.ServerConfig

	NacosConfig config.NacosConfig
)
