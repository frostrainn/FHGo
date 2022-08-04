package initialize

import (
	"fhgo/drivers"
	"fhgo/global"
	"fmt"
)

func ConnDB() {
	mysql, err := drivers.Mysql()
	if err != nil {
		fmt.Println("连接错误")
		return
	}
	fmt.Println("连接成功")
	fmt.Println(mysql)
	global.DB = mysql
}
