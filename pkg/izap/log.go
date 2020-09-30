package izap

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"learn_go/config"
	"learn_go/utils"
	"os"
	"sync"
	"time"
)

var debugLogger *zap.SugaredLogger = nil //
var serverLogger *zap.SugaredLogger = nil
var errorLogger *zap.SugaredLogger = nil //
var downloadPkgLogger *zap.Logger = nil  //

var (
	serverOnce sync.Once
	debugOnce  sync.Once
	errorOnce  sync.Once
	separator  = string(os.PathSeparator)
)

var zapEventLogConfig = zapcore.EncoderConfig{
	//Keys can be anything except the empty string.
	TimeKey:    "T",
	MessageKey: "M",
	LineEnding: zapcore.DefaultLineEnding,
	//EncodeLevel:    zapcore.CapitalLevelEncoder,
	EncodeTime:     zapcore.ISO8601TimeEncoder,
	EncodeDuration: zapcore.StringDurationEncoder,
	//EncodeCaller:   zapcore.ShortCallerEncoder,
}

var zapBigDataLogConfig = zapcore.EncoderConfig{
	//Keys can be anything except the empty string.
	MessageKey: "M",
	LineEnding: zapcore.DefaultLineEnding,
	//EncodeLevel:    zapcore.CapitalLevelEncoder,
	//EncodeTime:     zapcore.ISO8601TimeEncoder,
	EncodeDuration: zapcore.StringDurationEncoder,
	//EncodeCaller:   zapcore.ShortCallerEncoder,
}

func Init() {
	Get()
	GetDebug()
	GetError()
	newDownloadPkgLogger()
}

func Get() *zap.SugaredLogger {
	if serverLogger == nil {
		newSugarLogger()
	}
	return serverLogger
}

func GetDebug() *zap.SugaredLogger {
	if debugLogger == nil {
		newDebugLogger()
	}
	return debugLogger
}

func GetError() *zap.SugaredLogger {
	if errorLogger == nil {
		newErrorLogger()
	}
	return errorLogger
}

// --------------------------------------------------

func newSugarLogger() *zap.SugaredLogger {
	serverOnce.Do(func() {
		logDir := fmt.Sprintf("%s%sserver-log", config.Cfg.LogDir, separator)
		zapCnf := zap.NewDevelopmentEncoderConfig()
		l := doCreate(config.Cfg.Env, logDir, "server.log", zapCnf, 24*time.Hour)
		serverLogger = l.Sugar()
	})

	return serverLogger
}

func newDebugLogger() *zap.SugaredLogger {
	debugOnce.Do(func() {
		logDir := fmt.Sprintf("%s%sdebug-log", config.Cfg.LogDir, separator)
		zapCnf := zap.NewDevelopmentEncoderConfig()
		l := doCreate(config.Cfg.Env, logDir, "debug.log", zapCnf, 24*time.Hour)
		debugLogger = l.Sugar()
	})
	return debugLogger
}

func newErrorLogger() *zap.SugaredLogger {
	errorOnce.Do(func() {
		logDir := fmt.Sprintf("%s%serror-log", config.Cfg.LogDir, separator)
		zapCnf := zap.NewDevelopmentEncoderConfig()
		l := doCreate(config.Cfg.Env, logDir, "err.log", zapCnf, 24*time.Hour)
		errorLogger = l.Sugar()
	})

	return errorLogger
}

func newDownloadPkgLogger() {
	logDir := fmt.Sprintf("%s%sdownload-pkg", config.Cfg.LogDir, separator)
	l := doCreate(config.Cfg.Env, logDir, "pkg.log", zapEventLogConfig, 24*time.Hour)
	downloadPkgLogger = l
}

func doCreate(profile config.Env, logDir, logName string, cnf zapcore.EncoderConfig, rotationTime time.Duration) *zap.Logger {

	var logger *zap.Logger
	if profile == config.Production {
		utils.EnsureDir(logDir)
		// 创建 logger
		writer := getWriter(logDir+separator+logName, rotationTime)
		ws := zapcore.AddSync(writer)

		encoder := zapcore.NewConsoleEncoder(cnf)
		core := zapcore.NewCore(
			encoder,
			ws,
			zap.InfoLevel,
		)

		logger = zap.New(core)
	} else {
		var cfg zap.Config
		cfg = zap.NewDevelopmentConfig()
		cfg.EncoderConfig = zap.NewDevelopmentEncoderConfig()
		l, err := cfg.Build()
		if err != nil {
			panic(err)
		}
		logger = l

	}

	return logger
}

func getWriter(filename string, rotationTime time.Duration) io.Writer {
	// 生成rotatelogs的Logger 实际生成的文件名 demo.log.YYmmddHH
	// demo.log是指向最新日志的链接
	// 保存7天内的日志，每1小时(整点)分割一次日志
	format := ""
	if rotationTime > 10*time.Minute {
		format = ".%Y%m%d%H"
	} else {
		format = ".%Y%m%d%H%M"
	}
	hook, err := rotatelogs.New(
		filename+format, // 没有使用go风格反人类的format格式
		rotatelogs.WithLinkName(filename),
		rotatelogs.WithMaxAge(time.Hour*24*7),
		rotatelogs.WithRotationTime(rotationTime),
	)

	if err != nil {
		panic(err)
	}
	return hook
}
