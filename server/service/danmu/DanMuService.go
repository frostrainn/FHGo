package danmu

import "github.com/gin-gonic/gin"

type PostAndGetService struct {
}

func (p *PostAndGetService) PostDanMu(c *gin.Context) {

}

func (p *PostAndGetService) GetDanMu(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
