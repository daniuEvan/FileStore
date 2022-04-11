/**
 * @date: 2022/4/11
 * @desc: ...
 */

package router

import (
	"FileStore/global"
	"FileStore/middleware"
	"FileStore/service/controller/uploadController"
	"github.com/gin-gonic/gin"
)

func InitUploadRouter(routerGroup *gin.RouterGroup) {
	uploadRouter := routerGroup.Group("upload").Use(middleware.GinLogger(global.Logger))
	{
		uploadRouter.POST("/file/", uploadController.UploadFile) // 文件上传
	}
}
