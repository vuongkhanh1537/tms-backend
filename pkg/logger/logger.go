package logger

import (
	"os"

	"github.com/natefinch/lumberjack"
	"github.com/vuongkhanh1537/tms-backend/pkg/setting"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerZap struct {
	*zap.Logger
}

func NewLogger(config setting.LogConfig) *LoggerZap {
	level := getLogLevel(config.LogLevel)
	encoder := getEncoderLog()
	hook := lumberjack.Logger{
		Filename: config.FileLogName,
		MaxSize: config.MaxSize,
		MaxAge: config.MaxAge,
		MaxBackups: config.MaxBackups,
		Compress: config.Compress,
	}
	core := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)),
		level,
	)

	return &LoggerZap{
		zap.New(core, zap.AddStacktrace(zap.ErrorLevel)),
	}
}

func getLogLevel(log_level string) zapcore.Level {
	switch log_level {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "warn":
		return zap.WarnLevel
	case "error":
		return zap.ErrorLevel
	default:
		return zap.InfoLevel
	}
}

func getEncoderLog() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encodeConfig.TimeKey = "time"
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encodeConfig.EncodeCaller = zapcore.FullCallerEncoder

	return zapcore.NewConsoleEncoder(encodeConfig)
}