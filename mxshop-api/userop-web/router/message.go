package router

import (
	"github.com/gin-gonic/gin"
	"xfront/userop-web/api/message"
	"xfront/userop-web/middlewares"
)

func InitMessageRouter(Router *gin.RouterGroup) {
	MessageRouter := Router.Group("message").Use(middlewares.JWTAuth())
	{
		MessageRouter.GET("", message.List) // 轮播图列表页
		MessageRouter.POST("", message.New) //新建轮播图
	}
}
