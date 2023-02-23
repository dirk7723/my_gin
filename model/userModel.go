package model

import (
	"gorm.io/gorm"
	"shutuiche.com/luka/go_test/global"
)

type UserModel struct{}

type User struct {
	Id               int    `json:"id"`
	Truename         string `json:"truename"`
	Nickname         string `json:"nickname"`
	Kol_avatar_url   string `json:"kol_avatar_url"`
	Class_id         int    `json:"class_id"`
	Class_name       string `json:"class_name"`
	Fans_nums        int    `json:"fans_nums"`
	Zan_collent_nums int    `json:"zan_collent_nums"`
	Interactive_nums int    `json:"interactive_nums"`
	Leave_date       string `json:"leave_date"`
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
