package global

import (
	ut "github.com/go-playground/universal-translator"
	"xfront/goods-web/config"
	"xfront/goods-web/proto"
)

var (
	Trans ut.Translator

	ServerConfig *config.ServerConfig = &config.ServerConfig{}

	NacosConfig *config.NacosConfig = &config.NacosConfig{}

	GoodsSrvClient proto.GoodsClient

	InventorySrvClient proto.InventoryClient
)
