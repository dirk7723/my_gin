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
	dsn := "shutuiche:shutuiche2020@tcp(121.5.244.122:3306)/michael_kors?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "mk_",
			SingularTable: true, //使用单数表名，启用该选项后，`User`表将是 `user`
			NameReplacer:  nil,
			NoLowerCase:   false,
		},
	})
	return Db, err
}
