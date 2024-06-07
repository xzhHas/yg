package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha" // Go语言的图形验证码库
	"go.uber.org/zap"
	"net/http"
)

var store = base64Captcha.DefaultMemStore

func GetCaptcha(ctx *gin.Context) {
	// 图形验证码
	// 宽度、高度、验证码字符数、验证码噪点数、验证码宽度
	driver := base64Captcha.NewDriverDigit(80, 240, 5, 0.2, 80)
	cp := base64Captcha.NewCaptcha(driver, store)
	id, b64s, _, err := cp.Generate()
	if err != nil {
		zap.S().Errorf("生成验证码错误,: ", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "生成验证码错误",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"captchaId": id,
		"picPath":   b64s,
	})
}
