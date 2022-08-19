package system

import "github.com/gin-gonic/gin"

type Hello struct {
}

func (h *Hello) Hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
