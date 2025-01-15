package mysql

import (
	"errors"
	"forum/config"
	"forum/global"
	"forum/logger"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDB() (err error) {
	dsn := config.AppConf.MySQLConfig.Dsn
	Db, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		zap.L().Error("connect database failed")
		logger.Error("db connect error" + err.Error())
	} else {
		logger.Info("db connect success")
	}
	sqlDb, err := Db.DB()
	if err != nil {
		zap.L().Error("connect database failed")
		logger.Error("db connect error" + err.Error())
		return err
	}
	sqlDb.SetMaxIdleConns(config.AppConf.MySQLConfig.MaxIdleConns)
	sqlDb.SetMaxOpenConns(config.AppConf.MySQLConfig.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Duration(config.AppConf.MySQLConfig.MaxOpenConns))
	global.Db = Db
	return
}

// 定义业务状态
var (
	ErrorUserExit      = "用户已存在"
	ErrorUserNotExit   = "用户不已存在"
	ErrorPasswordWrong = "密码错误"
	ErrorGenIDFailed   = errors.New("创建用户ID失败")
	ErrorInvalidID     = "无效的ID"
	ErrorQueryFailed   = "查询数据失败"
	ErrorInsertFailed  = errors.New("插入数据失败")
)
