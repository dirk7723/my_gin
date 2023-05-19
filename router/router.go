package router

import (
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"shutuiche.com/luka/go_test/controller"
	"shutuiche.com/luka/go_test/global"
	"shutuiche.com/luka/go_test/middleware"
	"shutuiche.com/luka/go_test/pkg/result"
)

func Router() *gin.Engine {
	router := gin.Default()
	//处理异常
	router.NoRoute(HandleNotFound)
	router.NoMethod(HandleNotFound)
	router.Use(middleware.AccessLog())
	//跨域
	router.Use(middleware.Cors())
	router.Use(Recover)

	//static
	router.StaticFS("/static", http.Dir("."+global.StaticSetting.StaticDir))
	// 路径映射
	router.POST("/test/index", controller.TestController{}.Index)
	router.POST("/fsbot/index", controller.FsbotController{}.Index)

	FileController := controller.NewFileController()
	router.POST("/upload_file", FileController.UploadOne)
	lg := controller.NewLoginController()
	router.POST("/login", lg.Index)
	router.Any("/test", middleware.Auth(), func(c *gin.Context) {
		//request := c.MustGet("request").(string)
		c.JSON(http.StatusOK, gin.H{
			"msg": "success",
		})
	})
	router.POST("/menu", controller.Menucontroller{}.Index)
	return router
}

// 404
func HandleNotFound(c *gin.Context) {
	global.Logger.Errorf("handle not found: %v", c.Request.RequestURI)
	//global.Logger.Errorf("stack: %v",string(debug.Stack()))
	result.NewResult(c).Error(404, "资源未找到")
	return
}

// 500
func Recover(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			//打印错误堆栈信息
			//log.Printf("panic: %v\n", r)
			global.Logger.Errorf("panic: %v", r)
			//log stack
			global.Logger.Errorf("stack: %v", string(debug.Stack()))
			//print stack
			debug.PrintStack()
			//return
			result.NewResult(c).Error(500, "服务器内部错误")
		}
	}()
	//继续后续接口调用
	c.Next()
}
