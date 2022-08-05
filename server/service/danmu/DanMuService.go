package danmu

import (
	"fhgo/global"
	"fhgo/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PostAndGetService struct {
}

func (p *PostAndGetService) PostDanMu(c *gin.Context) {
	global.Logger.Info("SendDanMu")
	d := models.DanMu{}
	err := c.ShouldBind(&d)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "fail",
			"data":    nil,
		})
	}
	global.DB.Create(&d)
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "ok",
		"data":    nil,
	})
}

func (p *PostAndGetService) GetDanMu(c *gin.Context) {
	global.Logger.Info("GetDanMu")
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
