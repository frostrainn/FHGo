package initialize

import (
	"fhgo/service"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
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
