package logger

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/setting"
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"log"
	"os"
)

type tool struct {
	logger *logrus.Logger
}

func New(setting setting.Logger) Tool {
	logger := Logger()
	if setting.GetRunMode() != "debug" {
		logger = WriterLogger(setting)
	}
	return &tool{logger: logger}
}

func (l *tool) Trace(ctx *gin.Context, msg string) {
	if ctx != nil {
		l.logger.WithFields(GetFields(ctx, msg)).Trace(msg)
		return
	}
	l.logger.Trace(msg)
}

func (l *tool) Debug(ctx *gin.Context, msg string) {
	if ctx != nil {
		l.logger.WithFields(GetFields(ctx, msg)).Debug(msg)
		return
	}
	l.logger.Debug(msg)
}

func (l *tool) Info(ctx *gin.Context, msg string) {
	if ctx != nil {
		l.logger.WithFields(GetFields(ctx, msg)).Info(msg)
		return
	}
	l.logger.Info(msg)
}

func (l *tool) Warn(ctx *gin.Context, msg string) {
	if ctx != nil {
		l.logger.WithFields(GetFields(ctx, msg)).Warn(msg)
		return
	}
	l.logger.Warn(msg)
}

func (l *tool) Error(ctx *gin.Context, msg string) {
	logger := l.logger
	if ctx != nil {
		logger.WithFields(GetFields(ctx, msg)).Error(msg)
		return
	}
	logger.Error(msg)
}

func (l *tool) Fatal(ctx *gin.Context, msg string) {
	if ctx != nil {
		l.logger.WithFields(GetFields(ctx, msg)).Fatal(msg)
		return
	}
	l.logger.Fatal(msg)
}

func (l *tool) Panic(ctx *gin.Context, msg string) {
	if ctx != nil {
		l.logger.WithFields(GetFields(ctx, msg)).Panic(msg)
		return
	}
	l.logger.Panic(msg)
}

func Logger() *logrus.Logger {
	printLog := logrus.New()
	printLog.SetFormatter(&logrus.JSONFormatter{})
	printLog.SetOutput(os.Stdout)
	return printLog
}

func WriterLogger(setting setting.Logger) *logrus.Logger {
	// 創建 rotatelogs
	path := setting.GetLogFilePath() + "/" + setting.GetLogFileName() + "." + setting.GetLogFileExt()
	writer, err := rotatelogs.New(
		path+".%Y%m%d%H%M",                                        //分割後log文件名稱
		rotatelogs.WithLinkName(path),                             //當前log文件名稱
		rotatelogs.WithMaxAge(setting.GetLogMaxAge()),             //log文件存活時間
		rotatelogs.WithRotationTime(setting.GetLogRotationTime()), //切分log時間間隔
	)
	if err != nil {
		log.Fatalf(err.Error())
	}
	// 創建 hook
	hook := lfshook.NewHook(lfshook.WriterMap{
		logrus.TraceLevel: writer,
		logrus.DebugLevel: writer,
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, &logrus.JSONFormatter{})

	// 創建 logrus
	writeLogger := logrus.New()
	writeLogger.AddHook(hook)

	return writeLogger
}

func GetFields(ctx *gin.Context, msg string) map[string]interface{} {
	var path string
	var hostIP string
	var body interface{}
	var rid interface{}
	if ctx != nil {
		path = ctx.FullPath()
		hostIP = ctx.ClientIP()
		body = ctx.Value("Body")
		rid = ctx.Value("X-Request-Id")
	}
	fields := map[string]interface{}{
		"msg":     msg,
		"path":    path,
		"host_ip": hostIP,
		"body":    body,
		"rid":     rid,
	}
	return fields
}
