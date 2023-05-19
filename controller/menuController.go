package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"shutuiche.com/luka/go_test/model"
	"shutuiche.com/luka/go_test/pkg/result"
)

type Menucontroller struct{}

func (a Menucontroller) Index(c *gin.Context) {
	result := result.NewResult(c)
	//ret := map[string]interface{}{}
	//model.Menu

	var m1 *model.Menu
	ret := model.GetTreeMenu()
	m1 = ret[0]
	fmt.Println(m1.Name)
	result.Success(ret)
}
