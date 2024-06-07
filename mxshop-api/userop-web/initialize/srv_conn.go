package initialize

import (
	"fmt"
	_ "github.com/mbobakov/grpc-consul-resolver" // It's important
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"xfront/userop-web/global"
	"xfront/userop-web/proto"
)

func InitSrvConn() {
	consulInfo := global.ServerConfig.ConsulInfo
	// 商品微服务
	goodsConn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulInfo.Host, consulInfo.Port, global.ServerConfig.GoodsSrvInfo.Name),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		zap.S().Fatal("[InitSrvConn] 连接 【商品服务失败】")
	}
	global.GoodsSrvClient = proto.NewGoodsClient(goodsConn)

	// 用户功能微服务
	userOpConn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulInfo.Host, consulInfo.Port, global.ServerConfig.UserOpSrvInfo.Name),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		zap.S().Fatal("[InitSrvConn] 连接 【用户操作服务失败】")
	}
	global.UserFavClient = proto.NewUserFavClient(userOpConn)
	global.MessageClient = proto.NewMessageClient(userOpConn)
	global.AddressClient = proto.NewAddressClient(userOpConn)
}
