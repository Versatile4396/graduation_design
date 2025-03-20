package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type AppConfig struct {
	*AppBaseConfig `mapstructure:"app"`
	*MySQLConfig   `mapstructure:"mysql"`
	*LogConfig     `mapstructure:"logger"`
}

type AppBaseConfig struct {
	Name     string `mapstructure:"name"`
	Port     string `mapstructure:"port"`
	ImageDir string `mapstructure:"image_dir"`
	VideoDir string `mapstructure:"video_dir"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}
type RedisConfig struct {
	Host         string `mapstructure:"host"`
	Password     string `mapstructure:"password"`
	Port         int    `mapstructure:"port"`
	DB           int    `mapstructure:"db"`
	PoolSize     int    `mapstructure:"pool_size"`
	MinIdleConns int    `mapstructure:"min_idle_conns"`
}
type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DbName       string `mapstructure:"db_name"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	Dsn          string `mapstructure:"dsn"`
}

var AppConf = &AppConfig{}

func InitConfig() error {
	// 读取配置文件
	viper.SetConfigFile("./config/config.yaml")
	// 读取环境变量
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		viper.Unmarshal(AppConf)
	})
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("read file err ==> %v", err.Error()))
	}
	if err := viper.Unmarshal(AppConf); err != nil {
		panic(fmt.Errorf("unmarshal err => %v", err.Error()))
	}
	return err
}
