/*
 * @date: 2021/12/16
 * @desc: ...
 */

package initialize

import (
	"FileStore/database"
	"FileStore/global"
	"go.uber.org/zap"
)

// InitDatabase 初始化数据库
func InitDatabase() error {
	err := database.InitDB()
	if err != nil {
		global.Logger.Error("初始化数据库连接异常:", zap.String("error", err.Error()))
		return err
	}

	// 初始化redis
	err = database.InitRedisPool()
	if err != nil {
		global.Logger.Error("初始化 redis 连接异常:", zap.String("error", err.Error()))
		return err
	}
	return nil
}
