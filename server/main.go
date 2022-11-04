package main

import (
	"fhgo/initialize"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	r := initialize.Routers()
	http.ListenAndServe(":5486", r)
	//err := r.Run(":8080")
	//err = endless.ListenAndServe(":8080", r)	//优雅重启
	//if err != nil {
	//	return
	//}
}

func init() {
	initialize.InitLogger()
	initialize.InitConfig()  //读取Config文件
	initialize.ConnDB()      //连接数据库
	initialize.CreateTable() //创建表
	initialize.ConnRedis()   //连接Redis
	//pprof 服务器
	// /debug/pprof
	go func() {
		err := http.ListenAndServe(":6060", nil)
		if err != nil {
			return
		}
	}()
}
