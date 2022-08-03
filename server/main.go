package main

import (
	"fhgo/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	u := service.Group.UserService
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/", u.Login)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
