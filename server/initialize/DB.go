package initialize

import (
	"fhgo/drivers"
	"fhgo/global"
	"fhgo/models"
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

func CheckTableExists(dst ...interface{}) {
	err := global.DB.AutoMigrate(dst)
	if err != nil {
		return
	}
}

func CreateTable() {
	m := global.DB.Migrator()
	err := m.CreateTable(&models.DanMu{})
	if err != nil {
		return
	}

}
