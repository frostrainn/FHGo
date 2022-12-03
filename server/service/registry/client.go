package registry

import (
	"encoding/json"
	"google.golang.org/grpc/resolver"
	"io"
	"log"
	"net/http"
	"time"
)

type Producer struct {
	Registry  string
	NameSpace string
	Path      string
	Addr      string
	Duration  time.Duration
}

func NewProducer(Registry, NameSpace, Path, Addr string, Duration time.Duration) *Producer {
	Producer := &Producer{Registry: Registry, NameSpace: NameSpace, Path: Path, Addr: Addr, Duration: Duration}
	Producer.Heartbeat()
	return Producer
}

func (c *Producer) Heartbeat() {
	if c.Duration == 0 {
		c.Duration = defaultTimeout - defaultTimeout/5
	}
	var err error
	err = sendHeartbeat(c.Registry, c.NameSpace, c.Path, c.Addr)
	go func() {
		t := time.NewTicker(c.Duration)
		for err == nil {
			<-t.C
			err = sendHeartbeat(c.Registry, c.NameSpace, c.Path, c.Addr)
		}
	}()

}

func sendHeartbeat(registry, nameSpace, path, addr string) error {
	log.Println(addr, "send heart beat to registry", registry)
	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", registry, nil)
	if err != nil {
		panic(err) // NewRequest only errors on bad methods or un-parsable urls
	}
	req.Header.Set("X-rpc-Path", path)
	req.Header.Set("X-rpc-Addr", addr)
	req.Header.Set("X-rpc-NameSpace", nameSpace)
	if _, err := httpClient.Do(req); err != nil {
		log.Println("rpc server: heart beat err:", err)
		return err
	}
	return nil
}

type Consumer struct {
	Registry  string
	NameSpace string
	Api       string
	cc        resolver.ClientConn
	addrs     map[string][]string //服务列表

}

func NewConsumer(Registry, NameSpace, Path string) *Consumer {
	Consumer := &Consumer{Registry: Registry, NameSpace: NameSpace, Api: Path}
	Consumer.GetServers()
	return Consumer
}

func (c *Consumer) GetServers() []string {
	addrs := aliveServers(c.Registry, c.NameSpace, c.Api)
	c.addrs = make(map[string][]string)
	c.addrs[c.Api] = addrs
	return addrs
}

func (c *Consumer) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	paths := c.addrs[target.URL.Host]
	addrs := make([]resolver.Address, len(paths))
	for i, path := range paths {
		addrs[i] = resolver.Address{Addr: path}
	}
	cc.UpdateState(resolver.State{Addresses: addrs})
	return c, nil
}
func (c *Consumer) Scheme() string {
	return "zhubi"
}

// ResolveNow 监视目标更新
func (c *Consumer) ResolveNow(rn resolver.ResolveNowOptions) {
	log.Println("ResolveNow")
}

// Close 关闭
func (c *Consumer) Close() {
	log.Println("Close")
}

func aliveServers(registry, nameSpace, path string) []string {
	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", registry, nil)
	if err != nil {
		panic(err) // NewRequest only errors on bad methods or un-parsable urls
	}
	req.Header.Set("X-rpc-Path", path)
	req.Header.Set("X-rpc-NameSpace", nameSpace)
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Println("rpc server: heart beat err:", err)
		return nil
	}
	m := T{}
	body, _ := io.ReadAll(resp.Body)
	err = json.Unmarshal(body, &m)
	if err != nil {
		panic(err)
	}
	return m.XRpcAddr
}

type T struct {
	XRpcAddr []string `json:"X-rpc-Addr"`
}
