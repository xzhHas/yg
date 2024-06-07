package global

import (
	"gorm.io/gorm"
	"xback/order_srv/config"
	"xback/order_srv/proto"
)

var (
	DB *gorm.DB

	ServerConfig config.ServerConfig

	NacosConfig config.NacosConfig

	GoodsSrvClient proto.GoodsClient

	InventorySrvClient proto.InventoryClient
)
