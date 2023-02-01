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
	if s.mode == "release" {
		return s.vp.GetString("AppleLogin.Release.KeyName")
	}
	if s.mode == "production" {
		return s.vp.GetString("AppleLogin.Production.KeyName")
	}
	return ""
}

func (s *setting) GetBundleID() string {
	if s.mode == "debug" {
		return s.vp.GetString("AppleLogin.Debug.BundleID")
	}
	if s.mode == "release" {
		return s.vp.GetString("AppleLogin.Release.BundleID")
	}
	if s.mode == "production" {
		return s.vp.GetString("AppleLogin.Production.BundleID")
	}
	return ""
}

func (s *setting) GetDebugTokenURL() string {
	if s.mode == "debug" {
		return s.vp.GetString("AppleLogin.Debug.DebugTokenURL")
	}
	if s.mode == "release" {
		return s.vp.GetString("AppleLogin.Release.DebugTokenURL")
	}
	if s.mode == "production" {
		return s.vp.GetString("AppleLogin.Production.DebugTokenURL")
	}
	return ""
}

func (s *setting) GetTeamID() string {
	if s.mode == "debug" {
		return s.vp.GetString("AppleLogin.Debug.TeamID")
	}
	if s.mode == "release" {
		return s.vp.GetString("AppleLogin.Release.TeamID")
	}
	if s.mode == "production" {
		return s.vp.GetString("AppleLogin.Production.TeamID")
	}
	return ""
}

func (s *setting) GetKeyID() string {
	if s.mode == "debug" {
		return s.vp.GetString("AppleLogin.Debug.KeyID")
	}
	if s.mode == "release" {
		return s.vp.GetString("AppleLogin.Release.KeyID")
	}
	if s.mode == "production" {
		return s.vp.GetString("AppleLogin.Production.KeyID")
	}
	return ""
}
