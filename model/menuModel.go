package model

import "shutuiche.com/luka/go_test/global"

type MenuModel struct{}

type Menu struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Pid  int    `json:"pid"`
}

func GetTreeMenu() []*Menu {
	var menu []*Menu
	global.Db.Find(&menu)
	return menu
}
