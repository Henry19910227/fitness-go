package logger

import (
	"github.com/Henry19910227/fitness-go/config"
	"github.com/Henry19910227/fitness-go/internal/pkg/build"
	"github.com/prometheus/common/log"
	"github.com/spf13/viper"
	"time"
)

type setting struct {
	vp   *viper.Viper
	mode string
}

func New() Setting {
	path := config.RootPath() + "/config.yaml"
	vp := viper.New()
	vp.SetConfigFile(path)
	if err := vp.ReadInConfig(); err != nil {
		log.Fatalf(err.Error())
	}
	return &setting{vp: vp, mode: build.RunMode()}
}

func (s *setting) GetLogFilePath() string {
	if s.mode == "debug" {
		return s.vp.GetString("Log.Debug.Path")
	}
	if s.mode == "release" {
		return s.vp.GetString("Log.Release.Path")
	}
	if s.mode == "production" {
		return s.vp.GetString("Log.Production.Path")
	}
	return ""
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
	if s.mode == "release" {
		return s.vp.GetDuration("Log.Release.MaxAge") * time.Hour * 24
	}
	if s.mode == "production" {
		return s.vp.GetDuration("Log.Production.MaxAge") * time.Hour * 24
	}
	return 7 * time.Hour * 24
}

func (s *setting) GetLogRotationTime() time.Duration {
	if s.mode == "debug" {
		return s.vp.GetDuration("Log.Debug.RotationTime") * time.Minute
	}
	if s.mode == "release" {
		return s.vp.GetDuration("Log.Release.RotationTime") * time.Hour * 24
	}
	if s.mode == "production" {
		return s.vp.GetDuration("Log.Production.RotationTime") * time.Hour * 24
	}
	return 7 * time.Hour * 24
}

func (s *setting) GetRunMode() string {
	return s.mode
}
