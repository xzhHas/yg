package router

import (
	"github.com/gin-gonic/gin"
	"xfront/goods-web/middlewares"

	"xfront/goods-web/api/goods"
)

func InitGoodsRouter(Router *gin.RouterGroup) {
	GoodsRouter := Router.Group("goods").Use(middlewares.Trace())
	//GoodsRouter := Router.Group("goods")
	{
		GoodsRouter.GET("", goods.List)
		GoodsRouter.GET("/new", goods.NewHomeList)
		GoodsRouter.GET("/hot", goods.HotHomeList)
		GoodsRouter.GET("/search", goods.SearchGoods)
		// online前台主页面只显示，推荐的商品
		GoodsRouter.POST("", middlewares.JWTAuth(), middlewares.IsAdminAuth(), goods.New)          //改接口需要管理员权限
		GoodsRouter.GET("/:id", goods.Detail)                                                      //获取商品的详情
		GoodsRouter.DELETE("/:id", middlewares.JWTAuth(), middlewares.IsAdminAuth(), goods.Delete) //删除商品
		GoodsRouter.GET("/:id/stocks", goods.Stocks)                                               //获取商品的库存

		GoodsRouter.PUT("/:id", middlewares.JWTAuth(), middlewares.IsAdminAuth(), goods.Update)
		GoodsRouter.PATCH("/:id", middlewares.JWTAuth(), middlewares.IsAdminAuth(), goods.UpdateStatus)
	}
}
