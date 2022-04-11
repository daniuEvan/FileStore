/*
 * @date: 2021/12/15
 * @desc: ...
 */

package database

import (
	"FileStore/global"
	"FileStore/service/model"
	"context"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var serverDB *gorm.DB

// InitDB 初始化mysql
func InitDB() (err error) {
	db, err := newDBConn()
	if err != nil {
		global.Logger.Error("初始化数据库连接异常:", zap.String("errorMsg", err.Error()))
		return err
	}
	err = db.AutoMigrate(model.Models...)
	if err != nil {
		global.Logger.Error("数据库迁移失败:", zap.String("errorMsg", err.Error()))
		return err
	}
	serverDB = db
	return nil
}

// GetDB 获取数据库连接
func GetDB(ctx context.Context) (*gorm.DB, error) {
	// WithContext 实际是调用 db.Session(&Session{Context: ctx})，每次创建新 Session，各 db 操作之间互不影响
	dbManger, _ := serverDB.DB()
	err := dbManger.Ping()
	if err != nil {
		global.Logger.Warn("数据库连接异常,正在重新初始化数据库连接:", zap.String("error", err.Error()))
		err := dbManger.Close()
		if err != nil {
			global.Logger.Error("数据库连接关闭异常:", zap.String("error", err.Error()))
			return nil, err
		}
		err = InitDB()
		if err != nil {
			global.Logger.Error("初始化数据库连接异常:", zap.String("error", err.Error()))
			return nil, err
		}
	}
	return serverDB.WithContext(ctx), nil
}

// newDBConn 获取数据库连接
func newDBConn() (db *gorm.DB, err error) {
	dbType := global.ServerConfig.DatabaseInfo.DBType
	databasePoolStatus := global.ServerConfig.DatabaseInfo.OrmDatabasePoolInfo.Status
	tablePrefix := global.ServerConfig.DatabaseInfo.TablePrefix
	switch dbType {
	case "mysql":
		mysqlInfo := global.ServerConfig.DatabaseInfo.MysqlInfo
		dsn := fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			mysqlInfo.Username,
			mysqlInfo.Password,
			mysqlInfo.Host,
			mysqlInfo.Port,
			mysqlInfo.DBName,
		)
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix:   tablePrefix,
				SingularTable: true,
			},
		})
		if err != nil {
			global.Logger.Error("连接数据库异常:", zap.String("errorMsg", err.Error()))
			return nil, err
		}
	case "postgres":
		postgresInfo := global.ServerConfig.DatabaseInfo.PostgresInfo
		dsn := fmt.Sprintf(
			"user=%s password=%s host=%s  port=%d  dbname=%s search_path=%s sslmode=disable TimeZone=Asia/Shanghai",
			postgresInfo.Username,
			postgresInfo.Password,
			postgresInfo.Host,
			postgresInfo.Port,
			postgresInfo.DBName,
			postgresInfo.Schema,
		)
		db, err = gorm.Open(postgres.New(postgres.Config{
			DSN:                  dsn,
			PreferSimpleProtocol: true, // disables implicit prepared statement usage
		}), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix:   fmt.Sprintf("%s.%s", postgresInfo.Schema, tablePrefix), // pg 加上schema后 表前缀失效
				SingularTable: true,
			},
		})
		if err != nil {
			global.Logger.Error("连接数据库异常:", zap.String("errorMsg", err.Error()))
			return nil, err
		}
	default:
		global.Logger.Error("配置文件数据库dbType配置错误")
	}

	if databasePoolStatus == "disable" {
		return db, err
	}
	// 构建数据库连接池
	if err := buildOrmDatabasePool(); err != nil {
		return nil, err
	}
	return db, nil
}
