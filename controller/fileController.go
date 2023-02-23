package controller

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"shutuiche.com/luka/go_test/global"
	"shutuiche.com/luka/go_test/pkg/result"
)

type FileController struct{}

func NewFileController() FileController {
	return FileController{}
}

// 上传单张图片
func (a *FileController) UploadOne(c *gin.Context) {
	//save image
	f, err := c.FormFile("file")
	//错误处理
	result := result.NewResult(c)
	if err != nil {
		//fmt.Println(err.Error())
		result.Error(1006, "图片上传失败")
	} else {
		//将文件保存至本项目根目录中
		idstr := strconv.FormatInt(time.Now().Unix(), 10)
		destImage := global.ArticleImageSetting.UploadDir + "/" + idstr + ".jpg"
		err := c.SaveUploadedFile(f, "."+destImage)
		if err != nil {
			//fmt.Println("save err:")
			result.Error(1006, "图片保存失败")
		} else {
			imageUrl := global.ArticleImageSetting.ImageHost + destImage
			result.Success(gin.H{"url": imageUrl})
		}
	}
	//return
}
