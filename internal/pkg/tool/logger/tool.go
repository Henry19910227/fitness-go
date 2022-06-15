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
	RunMode string
	print   *logrus.Logger
	write   *logrus.Logger
}

func New(setting setting.Logger) Tool {
	writeLog, err := newWriteLogger(setting)
	if err != nil {
		log.Fatalf(err.Error())
	}
	printLog := newPrintLogger()
	runMode := setting.GetRunMode()
	return &tool{RunMode: runMode, print: printLog, write: writeLog}
}

func (l *tool) Trace(ctx *gin.Context, msg string) {
	fields := GetFields(ctx, msg)
	l.print.WithFields(fields).Trace(msg)
	if l.RunMode == "release" {
		l.write.WithFields(fields).Trace(msg)
	}
}

func (l *tool) Debug(ctx *gin.Context, msg string) {
	fields := GetFields(ctx, msg)
	l.print.WithFields(fields).Debug(msg)
	if l.RunMode == "release" {
		l.write.WithFields(fields).Debug(msg)
	}
}

func (l *tool) Info(ctx *gin.Context, msg string) {
	fields := GetFields(ctx, msg)
	l.print.WithFields(fields).Info(msg)
	if l.RunMode == "release" {
		l.write.WithFields(fields).Info(msg)
	}
}

func (l *tool) Warn(ctx *gin.Context, msg string) {
	fields := GetFields(ctx, msg)
	l.print.WithFields(fields).Warn(msg)
	if l.RunMode == "release" {
		l.write.WithFields(fields).Warn(msg)
	}
}

func (l *tool) Error(ctx *gin.Context, msg string) {
	fields := GetFields(ctx, msg)
	l.print.WithFields(fields).Error(msg)
	if l.RunMode == "release" {
		l.write.WithFields(fields).Error(msg)
	}
}

func (l *tool) Fatal(ctx *gin.Context, msg string) {
	fields := GetFields(ctx, msg)
	l.print.WithFields(fields).Fatal(msg)
	if l.RunMode == "release" {
		l.write.WithFields(fields).Fatal(msg)
	}
}

func (l *tool) Panic(ctx *gin.Context, msg string) {
	fields := GetFields(ctx, msg)
	l.print.WithFields(fields).Panic(msg)
	if l.RunMode == "release" {
		l.write.WithFields(fields).Panic(msg)
	}
}

func newPrintLogger() *logrus.Logger {
	printLog := logrus.New()
	printLog.SetFormatter(&logrus.JSONFormatter{})
	printLog.SetOutput(os.Stdout)
	return printLog
}

func newWriteLogger(setting setting.Logger) (*logrus.Logger, error) {
	// 創建 rotatelogs
	path := setting.GetLogFilePath()+"/"+setting.GetLogFileName()+"."+setting.GetLogFileExt()
	writer, err := rotatelogs.New(
		path+".%Y%m%d%H%M",                                        //分割後log文件名稱
		rotatelogs.WithLinkName(path),                             //當前log文件名稱
		rotatelogs.WithMaxAge(setting.GetLogMaxAge()),             //log文件存活時間
		rotatelogs.WithRotationTime(setting.GetLogRotationTime()), //切分log時間間隔
	)
	if err != nil {
		return nil, err
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

	return writeLogger, nil
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