package router

import (
	"github.com/gin-gonic/gin"
	"xfront/user-web/api"
	"xfront/user-web/middlewares"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	{
		UserRouter.GET("", middlewares.JWTAuth(), middlewares.IsAdminAuth(), api.GetUserList)
		// 登录，1.前台登录搭配captcha验证码一起使用 2.后台登录和前台使用方法一样
		UserRouter.POST("pwd_login", api.PassWordLogin)
		UserRouter.POST("register", api.Register)

		UserRouter.GET("detail", middlewares.JWTAuth(), api.GetUserDetail)
		UserRouter.PATCH("update", middlewares.JWTAuth(), api.UpdateUser)
	}
	//服务注册和发现
}
