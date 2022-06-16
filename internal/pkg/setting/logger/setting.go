package logger

import (
	"github.com/spf13/viper"
	"time"
)

type setting struct {
	vp *viper.Viper
	mode string
}

func New(vp *viper.Viper) Setting {
	return &setting{vp: vp, mode: vp.GetString("Server.RunMode")}
}

func (s *setting) GetLogFilePath() string {
	if s.mode == "debug" {
		return s.vp.GetString("Log.Debug.Path")
	}
	return s.vp.GetString("Log.Release.Path")
}

func (s *setting) GetLogFileName() string {
	return s.vp.GetString("Log.FileName")
}

func (s *setting) GetLogFileExt() string {
	return s.vp.GetString("Log.FileExt")
}

func (s *setting) GetLogMaxAge() time.Duration {
	if s.mode == "debug" {
		return s.vp.GetDuration("Log.Debug.MaxAge") * time.Minute
	}
	return s.vp.GetDuration("Log.Release.MaxAge") * time.Hour * 24
}

func (s *setting) GetLogRotationTime() time.Duration {
	if s.mode == "debug" {
		return s.vp.GetDuration("Log.Debug.RotationTime") * time.Minute
	}
	return s.vp.GetDuration("Log.Release.RotationTime") * time.Hour * 24
}

func (s *setting) GetRunMode() string {
	return s.vp.GetString("Server.RunMode")
}