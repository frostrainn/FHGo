package initialize

import (
	"fhgo/service"
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
		PublicGroup.POST("/SendDanMu", s.PostDanMu)
	}

	//PublicGroup.Handlers()
	//PublicGroup.GET()
	return Router
}
