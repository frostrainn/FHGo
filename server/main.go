package main

import (
	"fhgo/initialize"
	"github.com/fvbock/endless"
)

func main() {
	initialize.InitLogger()
	initialize.ConnDB()      //连接数据库
	initialize.CreateTable() //创建表
	initialize.ConnRedis()   //连接Redis
	r := initialize.Routers()
	err := endless.ListenAndServe(":8080", r)
	if err != nil {
		return
	}

}
