package logger

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var sugarLogger *zap.SugaredLogger

func InitLogger() {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	// logger := zap.New(core)
	logger := zap.New(core, zap.AddCaller())
	sugarLogger = logger.Sugar()
}
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	//如果想要追加写入可以查看我的博客文件操作那一章
	if _, err := os.Stat("./log"); os.IsNotExist(err) {
		os.Mkdir("./log", os.ModePerm)
	}
	fp := "./log/logger_%d.log"
	sub := 0
	for {
		if _, err := os.Stat(fmt.Sprintf(fp, sub)); os.IsNotExist(err) {
			break
		} else {
			sub++
		}
	}
	path := fmt.Sprintf(fp, sub)
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	return zapcore.AddSync(file)
}

func Debug(args ...interface{}) {
	sugarLogger.Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	sugarLogger.Debugf(template, args...)
}

func Info(args ...interface{}) {
	sugarLogger.Info(args...)
}

func Infof(template string, args ...interface{}) {
	sugarLogger.Infof(template, args...)
}

func Warn(args ...interface{}) {
	sugarLogger.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	sugarLogger.Warnf(template, args...)
}

func Error(args ...interface{}) {
	sugarLogger.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	sugarLogger.Errorf(template, args...)
}

func DPanic(args ...interface{}) {
	sugarLogger.DPanic(args...)
}

func DPanicf(template string, args ...interface{}) {
	sugarLogger.DPanicf(template, args...)
}

func Panic(args ...interface{}) {
	sugarLogger.Panic(args...)
}

func Panicf(template string, args ...interface{}) {
	sugarLogger.Panicf(template, args...)
}

func Fatal(args ...interface{}) {
	sugarLogger.Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	sugarLogger.Fatalf(template, args...)
}
