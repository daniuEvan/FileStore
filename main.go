/*
 * @date: 2021/12/14
 * @desc: ...
 */

package main

import (
	"FileStore/global"
	"FileStore/initialize"
	"fmt"
)

func main() {
	// 初始化配置文件
	initialize.InitConfigFromYaml()
	// 初始化 zap-logger
	initialize.InitLogger()
	// 初始化翻译
	if err := initialize.InitTrans(global.ServerConfig.Language); err != nil {
		global.Logger.Error(fmt.Sprintf("初始化翻译失败:%s", err.Error()))
		return
	}
	// 初始化自定义校验器
	initialize.InitCustomValidator()
	// 初始化数据库
	if err := initialize.InitDatabase(); err != nil {
		global.Logger.Error(fmt.Sprintf("初始化数据库失败:%s", err.Error()))
		return
	}
	// 初始化路由
	ginEngine := initialize.InitRouters()
	serverAddr := fmt.Sprintf("%s:%d", global.ServerConfig.Host, global.ServerConfig.Port)
	global.Logger.Info(fmt.Sprintf("handler running at %s", serverAddr))
	err := ginEngine.Run(serverAddr)
	if err != nil {
		global.Logger.Error(fmt.Sprintf("项目启动失败: %s", err.Error()))
	}
}
