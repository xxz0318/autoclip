// Package log 请修改此处包名注释
// @author: xiexinzhong
// @create: 2024-01-26 18:39
// @description:
package log

import (
	"io"
	"os"
	"time"

	"douyin_video/conf"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var SugarLogger *zap.SugaredLogger

func Debug(args ...interface{}) {
	SugarLogger.Debug(args...)
}

func Debugf(format string, args ...interface{}) {
	SugarLogger.Debugf(format, args...)
}

func Info(args ...interface{}) {
	SugarLogger.Info(args...)
}

func Infof(format string, args ...interface{}) {
	SugarLogger.Infof(format, args...)
}

func Warn(args ...interface{}) {
	SugarLogger.Warn(args...)
}

func Warnf(format string, args ...interface{}) {
	SugarLogger.Warnf(format, args...)
}

func Error(args ...interface{}) {
	SugarLogger.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	SugarLogger.Errorf(format, args...)
}
func Fatal(args ...interface{}) {
	SugarLogger.Fatal(args...)
}
func Fatalf(format string, args ...interface{}) {
	SugarLogger.Fatalf(format, args...)
}

// InitLog 初始化日志

func InitLog() {

	logDir := conf.C.LogDir
	encoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		MessageKey:  "msg",
		LevelKey:    "level",
		EncodeLevel: zapcore.CapitalLevelEncoder,
		TimeKey:     "ts",
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		CallerKey:    "file",
		EncodeCaller: zapcore.ShortCallerEncoder,
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(int64(d) / 1000000)
		},
	})
	// 实现两个判断日志等级的interface
	debugLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.DebugLevel
	})
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.InfoLevel
	})
	warnLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.WarnLevel
	})
	errorLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})
	debugWriter := getWriter(logDir + "%Y%m%d/" + conf.DebugFile)
	infoWriter := getWriter(logDir + "%Y%m%d/" + conf.InfoFile)
	warnWriter := getWriter(logDir + "%Y%m%d/" + conf.WarnFile)
	errorWriter := getWriter(logDir + "%Y%m%d/" + conf.ErrorFile)
	// 最后创建具体的Logger
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(debugWriter), debugLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(infoWriter), infoLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(warnWriter), warnLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(errorWriter), errorLevel),
		zapcore.NewCore(encoder, zapcore.Lock(os.Stdout), zap.DebugLevel),
	)
	log := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	SugarLogger = log.Sugar()
	defer log.Sync()
}
func getWriter(filename string) io.Writer {
	// 生成rotatelogs的Logger 实际生成的文件名 demo.log.YYmmddHH
	// demo.log是指向最新日志的链接
	// 保存7天内的日志，每1小时(整点)分割一次日志
	hook, err := rotatelogs.New(
		filename,
		// rotatelogs.WithLinkName(filename),
		rotatelogs.WithMaxAge(time.Hour*24*7),
		// rotatelogs.WithRotationTime(time.Hour),
	)
	if err != nil {
		panic(err)
	}
	return hook
}
