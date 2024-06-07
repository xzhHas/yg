package global

import (
	"github.com/olivere/elastic/v7"
	"gorm.io/gorm"
	"xback/goods_srv/config"
)

var (
	DB *gorm.DB

	ServerConfig config.ServerConfig

	NacosConfig config.NacosConfig

	EsClient *elastic.Client
)
