package api

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"
	"xfront/user-web/forms"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"

	"xfront/user-web/global"
)

func GenerateSmsCode(width int) string {
	// 用于生成指定长度的短信验证码
	numeric := [9]byte{1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		_, _ = fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}

func SendSms(ctx *gin.Context) {
	// 函数用于处理发送短信验证码的请求
	sendSmsForm := forms.SendSmsForm{}
	if err := ctx.ShouldBind(&sendSmsForm); err != nil {
		HandleValidatorError(ctx, err)
		return
	}
	fmt.Println("短信服务发送短信请求")

	// 创建阿里云短信客户端
	client, err := dysmsapi.NewClientWithAccessKey("cn-beijing", global.ServerConfig.AliSmsInfo.ApiKey, global.ServerConfig.AliSmsInfo.ApiSecrect)
	if err != nil {
		panic(err)
	}

	// 生成短信验证码
	smsCode := GenerateSmsCode(4) // 调用 GenerateSmsCode 函数生成指定长度的短信验证码

	// 创建发送短信请求
	request := requests.NewCommonRequest()
	request.Method = "POST"
	request.Scheme = "https"
	request.Domain = "dysmsapi.aliyuncs.com"
	request.Version = "2017-05-25"
	request.ApiName = "SendSms"
	request.QueryParams["RegionId"] = "cn-beijing"
	request.QueryParams["PhoneNumbers"] = sendSmsForm.Mobile
	request.QueryParams["SignName"] = "席万里开发平台"
	request.QueryParams["TemplateCode"] = "SMS_467155085"
	request.QueryParams["TemplateParam"] = "{\"code\":" + smsCode + "}"

	// 发送短信请求并获取响应
	response, err := client.ProcessCommonRequest(request)
	fmt.Println(client.DoAction(request, response))
	if err != nil {
		fmt.Print(err.Error())
	}

	// 将验证码保存到 Redis 中
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", global.ServerConfig.RedisInfo.Host, global.ServerConfig.RedisInfo.Port), // 设置 Redis 地址
	})
	// 手机号做key，验证码做value，设置过期时间
	rdb.Set(context.Background(), sendSmsForm.Mobile, smsCode, time.Duration(global.ServerConfig.RedisInfo.Expire)*time.Second) // 将验证码保存到 Redis 中，设置过期时间

	ctx.JSON(http.StatusOK, gin.H{"msg": "发送成功"})
}
