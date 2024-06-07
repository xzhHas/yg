package router

import (
	"github.com/gin-gonic/gin"
	"xfront/order-web/middlewares"

	"xfront/order-web/api/order"
	"xfront/order-web/api/pay"
)

func InitOrderRouter(Router *gin.RouterGroup) {
	OrderRouter := Router.Group("orders").Use(middlewares.JWTAuth()).Use(middlewares.Trace()) // tracing 监控、跟踪
	{
		OrderRouter.GET("", order.List)       // 订单列表
		OrderRouter.POST("", order.New)       // 新建订单
		OrderRouter.GET("/:id", order.Detail) // 订单详情
	}
	PayRouter := Router.Group("pay")
	{
		PayRouter.POST("alipay/notify", pay.Notify)
	}
}
