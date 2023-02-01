package fb_login

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

func (s *setting) GetAppID() string {
	if s.mode == "debug" {
		return s.vp.GetString("FBLogin.Debug.AppID")
	}
	if s.mode == "release" {
		return s.vp.GetString("FBLogin.Release.AppID")
	}
	if s.mode == "production" {
		return s.vp.GetString("FBLogin.Production.AppID")
	}
	return ""
}

func (s *setting) GetAppSecret() string {
	if s.mode == "debug" {
		return s.vp.GetString("FBLogin.Debug.AppSecret")
	}
	if s.mode == "release" {
		return s.vp.GetString("FBLogin.Release.AppSecret")
	}
	if s.mode == "production" {
		return s.vp.GetString("FBLogin.Production.AppSecret")
	}
	return ""
}

func (s *setting) GetDebugTokenURL() string {
	if s.mode == "debug" {
		return s.vp.GetString("FBLogin.Debug.DebugTokenURL")
	}
	if s.mode == "release" {
		return s.vp.GetString("FBLogin.Release.DebugTokenURL")
	}
	if s.mode == "production" {
		return s.vp.GetString("FBLogin.Production.DebugTokenURL")
	}
	return ""
}
