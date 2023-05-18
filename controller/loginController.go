package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"shutuiche.com/luka/go_test/model"
	"shutuiche.com/luka/go_test/pkg/result"
	"shutuiche.com/luka/go_test/pkg/util"
)

type LoginController struct{}

type login_form struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func NewLoginController() LoginController {
	return LoginController{}
}

func (a LoginController) Index(c *gin.Context) {
	var param login_form
	err := c.ShouldBind(&param)
	result := result.NewResult(c)
	if err != nil {
		result.Error(1006, "请传入参数")
	} else {
		list := model.UserModel{}.Find(map[string]interface{}{
			"username": param.Username,
			"password": param.Password,
		})
		fmt.Println(list)
		if len(list) == 0 {
			//data = gin.H{}
			result.Error(3, "账号或密码错误")
		} else {
			token, _ := util.GetToken(param.Username)
			fmt.Println(param.Username)
			ret := map[string]interface{}{
				"token": token,
			}
			// map[string]interface{}{
			// 	"username": param.Username,
			// 	"password": param.Password,
			// }
			result.Success(ret)
		}

	}
}
