package iab

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
	return s.vp.GetString("IAB.URL")
}

func (s *setting) GetTokenURL() string {
	return s.vp.GetString("IAB.TokenURL")
}

func (s *setting) GetScope() string {
	return s.vp.GetString("IAB.Scope")
}

func (s *setting) GetExpire() time.Duration {
	return s.vp.GetDuration("IAB.Expire") * time.Minute
}

func (s *setting) GetPackageName() string {
	if s.mode == "debug" {
		return s.vp.GetString("IAB.Debug.PackageName")
	}
	return s.vp.GetString("IAB.Release.PackageName")
}

func (s *setting) GetKeyName() string {
	if s.mode == "debug" {
		return s.vp.GetString("IAB.Debug.KeyName")
	}
	return s.vp.GetString("IAB.Release.KeyName")
}
