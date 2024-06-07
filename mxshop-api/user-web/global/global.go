package global

import (
	ut "github.com/go-playground/universal-translator"
	"xfront/user-web/config"
	"xfront/user-web/proto"
)

var (
	Trans ut.Translator

	ServerConfig *config.ServerConfig = &config.ServerConfig{}

	NacosConfig *config.NacosConfig = &config.NacosConfig{}

	UserSrvClient proto.UserClient
)
