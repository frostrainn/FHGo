package system

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"io"
)

type Hello struct {
}

func (h *Hello) Hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func (h *Hello) Sign(c *gin.Context) {

	body, _ := io.ReadAll(c.Request.Body)
	key := ""
	//for k,v :=range c.Request.Header {
	//	if k == "sign-key"{
	//		k = v[0]
	//	}
	//}
	key = c.Request.Header.Get("sign-key")
	decodeBytes, _ := base64.StdEncoding.DecodeString(key)
	b := hmacSha256(string(body), decodeBytes)
	sign := base64.StdEncoding.EncodeToString(b)

	c.JSON(200, gin.H{
		"sign": sign,
	})
}

func hmacSha256(data string, secret []byte) []byte {
	h := hmac.New(sha256.New, secret)
	h.Write([]byte(data))
	return h.Sum(nil)
}
