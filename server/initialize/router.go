package initialize

import (
	"fhgo/service"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	s := service.ServiceGroupApi.SystemService.ExampleService
	//
	PublicGroup := Router.Group("")
	{
		PublicGroup.GET("", s.Example)
		PublicGroup.POST("/SendDanMu")
	}

	//PublicGroup.Handlers()
	//PublicGroup.GET()
	return Router
}
