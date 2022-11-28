package fcm

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/build"
	"github.com/Henry19910227/fitness-go/internal/pkg/vp"
	"github.com/spf13/viper"
	"time"
)

type setting struct {
	vp   *viper.Viper
	mode string
}

func New() Setting {
	return &setting{vp: vp.Shared(), mode: build.RunMode()}
}

func (s *setting) GetURL() string {
	return s.vp.GetString("FCM.URL")
}

func (s *setting) GetTokenURL() string {
	return s.vp.GetString("FCM.TokenURL")
}

func (s *setting) GetScope() string {
	return s.vp.GetString("FCM.Scope")
}

func (s *setting) GetExpire() time.Duration {
	return s.vp.GetDuration("FCM.Expire") * time.Minute
}

func (s *setting) GetProjectID() string {
	if s.mode == "debug" {
		return s.vp.GetString("FCM.Debug.ProjectID")
	}
	return s.vp.GetString("FCM.Release.ProjectID")
}

func (s *setting) GetKeyName() string {
	if s.mode == "debug" {
		return s.vp.GetString("FCM.Debug.KeyName")
	}
	return s.vp.GetString("FCM.Release.KeyName")
}
