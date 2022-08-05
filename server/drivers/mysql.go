package drivers

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// Mysql 连接Mysql
func Mysql() (*gorm.DB, error) {
	dsn := "zzy:zzy@123@tcp(zhangzhongyuan.cn:53306)/live?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               dsn,
		DefaultStringSize: 171,
	}), &gorm.Config{
		SkipDefaultTransaction:                   false, //是否跳过自动事务，不跳过以保证数据安全性
		DisableForeignKeyConstraintWhenMigrating: true,  //建表时是否建立外键约束，关闭以提高性能。建议代码内体现逻辑外键
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, //表名命名规则：取结构体名的单数
		},
	})
	return db, err
}
