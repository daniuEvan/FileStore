/*
 * @date: 2021/12/16
 * @desc: ...
 */

package database

import (
	"FileStore/global"
	"go.uber.org/zap"
	"time"
)

// 构建 orm 数据库连接池
func buildOrmDatabasePool() error {
	sqlDB, err := serverDB.DB()
	if err != nil {
		global.Logger.Error("构建数据库连接池异常:", zap.String("errorMsg", err.Error()))
		return err
	}
	databasePoolInfo := global.ServerConfig.DatabaseInfo.OrmDatabasePoolInfo
	sqlDB.SetMaxIdleConns(databasePoolInfo.MaxIdleConns)                                    // 空闲连接池中连接的最大数量
	sqlDB.SetMaxOpenConns(databasePoolInfo.MaxOpenConns)                                    // 数据库连接的最大数量
	sqlDB.SetConnMaxLifetime(time.Minute * time.Duration(databasePoolInfo.ConnMaxLifetime)) // 连接可复用的最大时间。
	return nil
}
