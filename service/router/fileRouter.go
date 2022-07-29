/**
 * @date: 2022/4/11
 * @desc: ...
 */

package router

import (
	"FileStore/global"
	"FileStore/middleware"
	"FileStore/service/controller/fileController"
	"github.com/gin-gonic/gin"
)

func InitFileRouter(routerGroup *gin.RouterGroup) {
	fileRouter := routerGroup.Group("file").Use(middleware.GinLogger(global.Logger))
	{
		fileRouter.POST("/file_upload/", fileController.UploadFileHandler) // 文件上传
		fileRouter.GET("/file_get/", fileController.GetFileMetaHandler)    // 文件元信息获取
		fileRouter.GET("/file_download/", fileController.DownloadHandler)  // 文件下载
		fileRouter.PUT("/file_update/", fileController.UpdateHandler)      // 文件更新
		fileRouter.DELETE("/file_delete/", fileController.DeleteHandler)   // 文件删除
	}
}
