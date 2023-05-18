package global

import (
	//"github.com/jinzhu/gorm"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var Db *gorm.DB

func SetupDB() (*gorm.DB, error) {
	var err error
	dsn := "root:root@tcp(127.0.0.1:3306)/newbee?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "la_",
			SingularTable: true, //使用单数表名，启用该选项后，`User`表将是 `user`
			NameReplacer:  nil,
			NoLowerCase:   false,
		},
	})
	return Db, err
}
