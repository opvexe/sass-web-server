package tools

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var NormalLogger *zap.Logger

//初始化
func InitWithLog(env string, file string) {
	var level zapcore.Level
	switch env {
	case "dev":
		level = zapcore.DebugLevel
	case "test":
		level = zapcore.DebugLevel
	case "prod":
		level = zapcore.InfoLevel
	default:
		level = zapcore.InfoLevel
	}
	NormalLogger = NewLogger(file, level, 20, 3, 7, true)
}
