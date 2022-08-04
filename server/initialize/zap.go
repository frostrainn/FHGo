package initialize

import (
	"fhgo/global"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLogger() {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	logger := zap.New(core)
	global.Logger = logger.Sugar()
}

func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./log/run.log",
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func getEncoder() zapcore.Encoder {
	e := zap.NewProductionEncoderConfig()
	e.EncodeTime = zapcore.ISO8601TimeEncoder
	e.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(e)
}
