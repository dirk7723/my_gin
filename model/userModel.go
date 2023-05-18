package model

import (
	"gorm.io/gorm"
	"shutuiche.com/luka/go_test/global"
)

type UserModel struct{}

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

// type Param_user struct {
// 	Shop_name string `form:"shop_name"`
// }

func (m UserModel) List(param map[string]interface{}) []*User {
	filed := "u.uid,u.truename,u.avatar_url,k.avatar_url kol_avatar_url,u.class_id,class_name,FROM_BASE64(nickname) nickname,fans_nums,zan_collent_nums interactive_nums,leave_date"
	queryDB := global.Db.Table("mk_user u").Joins("join mk_kol k on k.uid = u.uid").Select(filed)
	if shop_name, ok := param["shop_name"].(string); ok {
		//动态加载查询条件
		queryDB.Clauses(gorm.Expr("u.shop_name like ?", shop_name+"%"))
	}
	var Users []*User
	queryDB.Scan(&Users)
	return Users
}

func (m UserModel) Find(param map[string]interface{}) []*User {
	var user []*User
	f1, f2 := false, false
	if _, ok := param["username"].(string); ok {
		//global.Db.Where("username = ?", username).First(&user)
		f1 = true
	}
	if _, ok := param["password"].(string); ok {
		//global.Db.Where("username = ?", username).First(&user)
		f2 = true
	}
	if f1 && f2 {
		global.Db.Where("username = ? AND password = ?", param["username"], param["password"]).First(&user)
	}
	return user
}
