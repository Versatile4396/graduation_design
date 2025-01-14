package Log

import (
	"be/config"
	"fmt"
	"os"

	"github.com/natefinch/lumberjack"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var lg *zap.Logger

func Init(cfg *config.LogConfig, mode string) (err error) {
	writeSyncer := getLogWriter(cfg.Filename, cfg.MaxSize, cfg.MaxBackups, cfg.MaxAge)
	encoder := getEncoder()
	var l = new(zapcore.Level)
	err = l.UnmarshalText([]byte(cfg.Level))
	if err != nil {
		return
	}
	var core zapcore.Core
	if mode == "dev" {
		// 进入开发模式，日志输出到终端
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		core = zapcore.NewTee( // 多个输出
			zapcore.NewCore(encoder, writeSyncer, l),                                     // 往日志文件里面写
			zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel), // 终端输出
		)
	} else {
		core = zapcore.NewCore(encoder, writeSyncer, l)
	}
	lg = zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(lg)
	zap.L().Info("init logger success")
	return
}
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}
func InitLogger() {
}
func getLogWriter(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
	}
	return zapcore.AddSync(lumberJackLogger)
}

// \033 这是标记变换颜色的起始标记，这之后的[1;31;40m则是定义颜色，1表示代码的意义或者说是显示方式，31表示前景颜色，40则是背景颜色。
// 在这定义之后终端就会显示你设定的样式，如果只是要改变一行的样式则在结尾加入\033[0m表示恢复终端默认样式。

func Info(msg string) {
	fmt.Printf("\033[1;32;42m%s\033[0m\n\n", msg)
}
func Warning(msg string) {
	fmt.Printf("\033[1;33;40m%s\033[0m\n\n", msg)
}

func Error(msg string) {
	fmt.Printf("\033[1;31;41m%s\033[0m\n\n", msg)
}

func Fmt(info interface{}) {
	fmt.Printf("\033[1;36;40m%+v\033[0m\n\n", info)
}
