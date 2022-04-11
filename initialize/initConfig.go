/*
 * @date: 2021/12/15
 * @desc: ...
 */

package initialize

import (
	"FileStore/global"
	"FileStore/utils"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func InitConfigFromYaml() {
	logger, _ := zap.NewDevelopment()
	debugEnv := utils.IsDebugEnv()
	configFilePrefix := "config"
	configFileName := fmt.Sprintf("./%s-dev.yaml", configFilePrefix)
	if !debugEnv {
		configFileName = fmt.Sprintf("./%s-pro.yaml", configFilePrefix)
		logger, _ = zap.NewProduction()
	}
	sugar := logger.Sugar()
	v := viper.New()
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		sugar.Errorw("配置初始化失败", "errMsg", err.Error())
		panic(err.Error())
	}
	if err := v.Unmarshal(global.ServerConfig); err != nil {
		sugar.Errorw("配置初始化失败", "errMsg", err.Error())
		panic(err.Error())
	}
	sugar.Info("配置文件初始化完成.")

	// 监测配置文件变化
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		// 配置文件发生变更之后会调用的回调函数
		sugar.Info("配置文件重新初始化:")
		if err := v.ReadInConfig(); err != nil {
			sugar.Errorw("配置重新初始化失败", "errMsg", err.Error())
			panic(err.Error())
		}
		if err := v.Unmarshal(global.ServerConfig); err != nil {
			sugar.Errorw("配置重新初始化失败", "errMsg", err.Error())
			panic(err.Error())
		}
		sugar.Info("配置文件重新初始化完成.")
	})

}
