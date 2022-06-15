package logger

import "time"

type Setting interface {
	GetLogFilePath() string
	GetLogFileName() string
	GetLogFileExt() string
	GetLogMaxAge() time.Duration
	GetLogRotationTime() time.Duration
	GetRunMode() string
}
