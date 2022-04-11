/**
 * @date: 2022/4/11
 * @desc: ...
 */

package router

import (
	"FileStore/global"
	"FileStore/middleware"
	"FileStore/service/controller/indexController"
	"github.com/gin-gonic/gin"
)

func InitIndexRouter(routerGroup *gin.Engine) {
	indexRouter := routerGroup.Group("index").Use(middleware.GinLogger(global.Logger))
	{
		indexRouter.GET("/", indexController.Index)
	}
}
