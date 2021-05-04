package tool

import (
	"github.com/Henry19910227/fitness-go/internal/setting"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
)

type logger struct {
	RunMode string
	print   *logrus.Logger
	write   *logrus.Logger
}

func NewLogger(setting setting.Logger) (Logger, error) {
	writeLog, err := newWriteLogger(setting)
	if err != nil {
		return nil, err
	}
	printLog := newPrintLogger()
	runMode := setting.GetRunMode()

	return &logger{runMode, printLog, writeLog}, nil
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

// Trace implement Logger interface
func (logger *logger) Trace(fields map[string]interface{}, msg string) {
	logger.print.WithFields(fields).Trace(msg)
	if logger.RunMode == "release" {
		logger.write.WithFields(fields).Trace(msg)
	}
}

// Debug implement Logger interface
func (logger *logger) Debug(fields map[string]interface{}, msg string) {
	logger.print.WithFields(fields).Debug(msg)
	if logger.RunMode == "release" {
		logger.write.WithFields(fields).Debug(msg)
	}
}

// Info implement Logger interface
func (logger *logger) Info(fields map[string]interface{}, msg string) {
	logger.print.WithFields(fields).Info(msg)
	if logger.RunMode == "release" {
		logger.write.WithFields(fields).Info(msg)
	}
}

// Warn implement Logger interface
func (logger *logger) Warn(fields map[string]interface{}, msg string) {
	logger.print.WithFields(fields).Warn(msg)
	if logger.RunMode == "release" {
		logger.write.WithFields(fields).Warn(msg)
	}
}

// Error implement Logger interface
func (logger *logger) Error(fields map[string]interface{}, msg string) {
	logger.print.WithFields(fields).Error(msg)
	if logger.RunMode == "release" {
		logger.write.WithFields(fields).Error(msg)
	}
}

// Fatal implement Logger interface
func (logger *logger) Fatal(fields map[string]interface{}, msg string) {
	logger.print.WithFields(fields).Fatal(msg)
	if logger.RunMode == "release" {
		logger.write.WithFields(fields).Fatal(msg)
	}
}

// Panic implement Logger interface
func (logger *logger) Panic(fields map[string]interface{}, msg string) {
	logger.print.WithFields(fields).Panic(msg)
	if logger.RunMode == "release" {
		logger.write.WithFields(fields).Panic(msg)
	}
}
