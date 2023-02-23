package controller

import (
	"github.com/gin-gonic/gin"
	"shutuiche.com/luka/go_test/model"
	"shutuiche.com/luka/go_test/pkg/result"
)

type TestController struct{}

type Product struct {
	Id    uint   `json:"id"`
	Code  string `json:"code"`
	Price uint   `json:"price"`
}
type Shop struct {
	Id        int    `json:"id"`
	Shop_name string `json:"shop_name"`
}

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

type Param_user struct {
	Shop_name string `form:"shop_name"`
}

func (a TestController) Index(c *gin.Context) {
	var param Param_user
	err := c.ShouldBind(&param)
	result := result.NewResult(c)
	if err != nil {
		result.Error(1006, "请传入参数")
	} else {
		Users := model.UserModel{}.List(map[string]interface{}{
			"shop_name": param.Shop_name,
		})
		result.Success(Users)
	}
}
