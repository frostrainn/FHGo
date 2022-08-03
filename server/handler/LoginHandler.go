package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
	// TODO : login type  0: passwd  1: token and cookie (without passwd )
	// TODO : if login type == 0 , get userName and passwd ,check the db ,  if pass return token and cookie (save the token to redis with 1 day )
}

func LogOut() {
	// todo : delete the user's token in redis
}

func ForgetPasswd() {
	// todo: you know
}

func CheckToken() {
	// todo: you know
}
