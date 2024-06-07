package global

import (
	ut "github.com/go-playground/universal-translator"
	"xfront/userop-web/config"
	"xfront/userop-web/proto"
)

var (
	Trans ut.Translator

	ServerConfig *config.ServerConfig = &config.ServerConfig{}

	NacosConfig *config.NacosConfig = &config.NacosConfig{}

	GoodsSrvClient proto.GoodsClient

	MessageClient proto.MessageClient

	AddressClient proto.AddressClient

	UserFavClient proto.UserFavClient
)
