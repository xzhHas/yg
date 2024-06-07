package router

import (
	"github.com/gin-gonic/gin"
	"xfront/userop-web/api/address"
	"xfront/userop-web/middlewares"
)

func InitAddressRouter(Router *gin.RouterGroup) {
	AddressRouter := Router.Group("address").Use(middlewares.JWTAuth())
	{
		AddressRouter.GET("", address.List)
		AddressRouter.DELETE("/:id", address.Delete)
		AddressRouter.POST("", address.New)
		AddressRouter.PUT("/:id", address.Update)
	}
}
