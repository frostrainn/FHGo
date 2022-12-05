package registry

import (
	"github.com/gin-gonic/gin"
	"sync"
	"time"
)

type ZRegistry struct {
	timeout time.Duration
	mu      sync.Mutex // protect following
	servers map[string][]*ServerItem
}

type ServerItem struct {
	Addr  string
	start time.Time
}

func (r *ZRegistry) Registry(c *gin.Context) {
	path := c.GetHeader("X-rpc-Path")
	addr := c.GetHeader("X-rpc-Addr")

	if path == "" {
		c.JSON(400, gin.H{"error": "path is empty"})
		return
	}

	if addr == "" {
		c.JSON(200, gin.H{"X-rpc-Addr": r.aliveServers(path)})
		return
	}

	r.putServer(path, addr)
	c.JSON(200, gin.H{"X-rpc-Addr": r.aliveServers(path)})

}

const defaultTimeout = time.Minute

func New(timeout time.Duration) *ZRegistry {
	return &ZRegistry{
		servers: make(map[string][]*ServerItem),
		timeout: timeout,
	}
}

var DefaultGeeRegister = New(defaultTimeout)
