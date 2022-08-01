package apple_login

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/build"
	"github.com/Henry19910227/fitness-go/internal/pkg/vp"
	"github.com/spf13/viper"
)

type setting struct {
	vp   *viper.Viper
	mode string
}

func New() Setting {
	return &setting{vp: vp.Shared(), mode: build.RunMode()}
}

func (s *setting) GetKeyName() string {
	if s.mode == "debug" {
		return s.vp.GetString("AppleLogin.Debug.KeyName")
	}
	return s.vp.GetString("AppleLogin.Release.KeyName")
}

func (s *setting) GetBundleID() string {
	if s.mode == "debug" {
		return s.vp.GetString("AppleLogin.Debug.BundleID")
	}
	return s.vp.GetString("AppleLogin.Release.BundleID")
}

func (s *setting) GetDebugTokenURL() string {
	if s.mode == "debug" {
		return s.vp.GetString("AppleLogin.Debug.DebugTokenURL")
	}
	return s.vp.GetString("AppleLogin.Release.DebugTokenURL")
}

func (s *setting) GetTeamID() string {
	if s.mode == "debug" {
		return s.vp.GetString("AppleLogin.Debug.TeamID")
	}
	return s.vp.GetString("AppleLogin.Release.TeamID")
}

func (s *setting) GetKeyID() string {
	if s.mode == "debug" {
		return s.vp.GetString("AppleLogin.Debug.KeyID")
	}
	return s.vp.GetString("AppleLogin.Release.KeyID")
}
