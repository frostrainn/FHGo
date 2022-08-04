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
		fmt.Println("MySQL数据库连接错误")
		return
	}
	fmt.Println("数据库连接建立成功")
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
