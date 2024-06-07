package global

import (
	ut "github.com/go-playground/universal-translator"
	"xfront/order-web/config"
	"xfront/order-web/proto"
)

var (
	Trans ut.Translator

	ServerConfig *config.ServerConfig = &config.ServerConfig{}

	NacosConfig *config.NacosConfig = &config.NacosConfig{}

	GoodsSrvClient proto.GoodsClient

	OrderSrvClient proto.OrderClient

	InventorySrvClient proto.InventoryClient
)
