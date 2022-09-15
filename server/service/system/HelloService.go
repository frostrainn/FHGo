package system

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"math/rand"
	"time"
)

type Hello struct {
}

type Sub struct {
	Id       int    `json:"id"`
	UserKey  string `json:"userkey"`
	Uid      string `json:"uid"`
	Icon     string `json:"icon"`
	Nickname string `json:"nickname"`
	Deleted  int    `json:"deleted"`
}

type full struct {
	RoleInfo []Sub `json:"roleInfos"`
}

func (h *Hello) Hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func (h *Hello) Report(c *gin.Context) {
	tag := rand.Intn(2)
	if tag == 1 {
		time.Sleep(10 * time.Second)
	}

	var p full
	err := json.NewDecoder(c.Request.Body).Decode(&p)
	if err != nil {
		c.JSON(404, gin.H{
			"code":    "404",
			"message": err.Error(),
		})
		log.Println(err)
		log.Println(p)
		return

	}
	log.Println(p)
	c.JSON(200, gin.H{
		"code":    "0",
		"message": "ok",
		"obj":     p,
	})
}

func (h *Hello) Sign(c *gin.Context) {

	time.Sleep(time.Second * 20)
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
