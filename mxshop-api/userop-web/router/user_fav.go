package router

import (
	"github.com/gin-gonic/gin"
	"xfront/userop-web/api/user_fav"
	"xfront/userop-web/middlewares"
)

func InitUserFavRouter(Router *gin.RouterGroup) {
	UserFavRouter := Router.Group("userfavs").Use(middlewares.JWTAuth())
	{
		UserFavRouter.DELETE("/:id", user_fav.Delete) // 删除收藏记录
		UserFavRouter.GET("/:id", user_fav.Detail)    // 获取收藏记录
		UserFavRouter.POST("", user_fav.New)          // 新建收藏记录
		UserFavRouter.GET("", user_fav.List)          // 获取当前用户的收藏
	}
}
