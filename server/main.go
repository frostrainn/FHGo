package main

import (
	"fhgo/initialize"
)

func main() {
	initialize.InitLogger()
	initialize.ConnDB()      //连接数据库
	initialize.CreateTable() //创建表
	r := initialize.Routers()
	err := r.Run()
	if err != nil {
		return
	}
}
