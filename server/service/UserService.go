package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserService struct{}

func (u *UserService) Login(c *gin.Context) {
	{
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	}
}
