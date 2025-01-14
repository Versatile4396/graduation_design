package mysql

import (
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
