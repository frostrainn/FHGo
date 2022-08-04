package danmu

import (
	"fhgo/global"
	"github.com/gin-gonic/gin"
)

type PostAndGetService struct {
}

func (p *PostAndGetService) PostDanMu(c *gin.Context) {

}

func (p *PostAndGetService) GetDanMu(c *gin.Context) {
	global.Logger.Info("GetDanMu")
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
