/*
 * @date: 2021/12/15
 * @desc: ...
 */

package initialize

import (
	"FileStore/middleware"
	userRouter "FileStore/service/router"
	"FileStore/utils"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func InitRouters() *gin.Engine {
	var defaultEngine *gin.Engine
	if utils.IsDebugEnv() {
		defaultEngine = gin.Default()
		pprof.Register(defaultEngine)
	} else {
		gin.SetMode(gin.ReleaseMode)
		defaultEngine = gin.New()
	}
	defaultEngine.Use(middleware.Cors()) // 跨域
	apiGroup := defaultEngine.Group("api/v1")
	userRouter.InitUserRouter(apiGroup)
	return defaultEngine

}
