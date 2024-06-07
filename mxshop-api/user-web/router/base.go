package router

import (
	"github.com/gin-gonic/gin"
	"xfront/user-web/api"
)

func InitBaseRouter(Router *gin.RouterGroup) {
	BaseRouter := Router.Group("base")
	{
		// 验证码，1.管理页面的登陆 2.前台页面的登陆
		BaseRouter.GET("captcha", api.GetCaptcha)
		// 阿里云发送短信服务，1.前台页面注册的时候使用到了
		BaseRouter.POST("send_sms", api.SendSms)
	}
}
