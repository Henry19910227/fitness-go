package setting

import (
	"github.com/spf13/viper"
	"time"
)

type logSetting struct {
	vp   *viper.Viper
	mode string
}

func NewLogger(viperTool *viper.Viper) Logger {
	return &logSetting{viperTool, viperTool.GetString("Server.RunMode")}
}

func (setting *logSetting) GetLogMaxAge() time.Duration {
	if setting.mode == "debug" {
		return setting.vp.GetDuration("Log.Debug.MaxAge") * time.Minute
	}
	return setting.vp.GetDuration("Log.Release.MaxAge") * time.Hour * 24
}

func (setting *logSetting) GetLogRotationTime() time.Duration {
	if setting.mode == "debug" {
		return setting.vp.GetDuration("Log.Debug.RotationTime") * time.Minute
	}
	return setting.vp.GetDuration("Log.Release.RotationTime") * time.Hour * 24
}

func (setting *logSetting) GetLogFilePath() string {
	if setting.mode == "debug" {
		return setting.vp.GetString("Log.Debug.Path")
	}
	return setting.vp.GetString("Log.Release.Path")
}

func (setting *logSetting) GetLogFileName() string {
	return setting.vp.GetString("Log.FileName")
}

func (setting *logSetting) GetLogFileExt() string {
	return setting.vp.GetString("Log.FileExt")
}

func (setting *logSetting) GetRunMode() string {
	return setting.vp.GetString("Server.RunMode")
}