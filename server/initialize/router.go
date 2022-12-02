package initialize

import (
	"fhgo/service"
	"fhgo/service/registry"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	//console 日志染色 关闭/打开
	//gin.DisableConsoleColor()
	gin.ForceConsoleColor()

	Router := gin.Default()
	s := service.ServiceGroupApi.DanMuService.PostAndGetService
	//
	PublicGroup := Router.Group("")
	{
		PublicGroup.GET("/GetDanMu", s.GetDanMu)
		PublicGroup.GET("/getAny", s.GetAny)
		PublicGroup.POST("/SendDanMu", s.PostDanMu)
	}

	sys := service.ServiceGroupApi.SystemService.Hello
	SystemGroup := Router.Group("")
	{
		SystemGroup.GET("/hello", sys.Hello)
		SystemGroup.POST("/sign/vgHweKN", sys.Sign)
		SystemGroup.POST("/report", sys.Report)
	}

	reg := registry.DefaultGeeRegister

	RegisterGroup := Router.Group("/_rpc_/registry")
	{
		RegisterGroup.GET("registry", reg.Registry)
	}

	//PublicGroup.Handlers()
	//PublicGroup.GET()
	return Router
}
