package main

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"

	"xfront/user-web/utils/register/consul"

	"xfront/user-web/global"
	"xfront/user-web/initialize"
	myvalidator "xfront/user-web/validator"
)

func main() {
	//1. 初始化logger
	initialize.InitLogger()
	//2. 初始化配置文件
	initialize.InitConfig()
	//3. 初始化routers
	Router := initialize.Routers()
	//4. 初始化翻译
	if err := initialize.InitTrans("zh"); err != nil {
		panic(err)
	}
	//5. 初始化srv的连接
	initialize.InitSrvConn()

	viper.AutomaticEnv()
	// 我在这里统一使用我的测试端口，不适用自动获取端口了
	// goods:8021  order:8022 oss:8023  user:8024  userop:8025
	//如果是本地开发环境端口号固定，线上环境启动获取端口号
	//debug := viper.GetBool("MXSHOP_DEBUG")
	//if !debug {
	//	port, err := utils.GetFreePort()
	//	if err == nil {
	//		global.ServerConfig.Port = port
	//	}
	//}

	//注册验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("mobile", myvalidator.ValidateMobile)
		_ = v.RegisterTranslation(
			"mobile",
			global.Trans,
			func(ut ut.Translator) error {
				return ut.Add("mobile", "{0} 非法的手机号码!", true)
			},
			func(ut ut.Translator, fe validator.FieldError) string {
				t, _ := ut.T("mobile", fe.Field())
				return t
			})
	}

	//服务注册
	register_client := consul.NewRegistryClient(global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port)
	serviceId := fmt.Sprintf("%s", uuid.NewV4())
	fmt.Println(global.ServerConfig)
	err := register_client.Register(global.ServerConfig.Host, global.ServerConfig.Port, global.ServerConfig.Name, []string{"imooc", "bobby", "user", "srv"}, serviceId)
	if err != nil {
		zap.S().Panic("服务注册失败:", err.Error())
	}

	zap.S().Debugf("启动服务器, 端口： %d", global.ServerConfig.Port)
	if err := Router.Run(fmt.Sprintf(":%d", global.ServerConfig.Port)); err != nil {
		zap.S().Panic("启动失败:", err.Error())
	}

	//接收终止信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
