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
		global.Logger.Error("数据库连接失败")
		return
	}
	global.Logger.Info("数据库连接成功")
	global.DB = mysql
}

func ConnRedis() {
	drivers.Redis()
	fmt.Println("Redis连接成功")
}

func CheckTableExists(dst ...interface{}) {
	err := global.DB.AutoMigrate(dst)
	if err != nil {
		return
	}
}

func CreateTable() {
	m := global.DB.Migrator()
	if !m.HasTable(&models.DanMu{}) {
		global.Logger.Info("DanMu表不存在，新建DanMu表")
		err := m.CreateTable(&models.DanMu{})
		if err != nil {
			global.Logger.Error("新建DanMu表失败")
			return
		}
	}
}
