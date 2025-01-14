package config

import (
	"log/slog"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func initDB() (err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/gindb?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		slog.Error("db connect error", slog.String("error", err.Error()))
	}
	slog.Info("db connect success", slog.String("db", Db.Name()))
	return
}
