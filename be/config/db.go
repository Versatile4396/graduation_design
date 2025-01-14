package config

import (
	"log/slog"

	"time"

	"be/global"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func initDB() (err error) {
	dsn := AppConfig.Database.Dsn
	Db, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		slog.Error("db connect error", slog.String("error", err.Error()))
	}
	slog.Info("db connect success", slog.String("db", Db.Name()))

	sqlDb, err := Db.DB()
	if err != nil {
		slog.Error("db connect error", slog.String("error", err.Error()))
	}
	sqlDb.SetMaxIdleConns(AppConfig.Database.MaxIdleConns)
	sqlDb.SetMaxOpenConns(AppConfig.Database.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Duration(AppConfig.Database.MaxOpenConns))
	global.Db = Db
	return
}
