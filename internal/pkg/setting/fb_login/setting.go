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
	return s.vp.GetString("FBLogin.Release.AppID")
}

func (s *setting) GetAppSecret() string {
	if s.mode == "debug" {
		return s.vp.GetString("FBLogin.Debug.AppSecret")
	}
	return s.vp.GetString("FBLogin.Release.AppSecret")
}

func (s *setting) GetDebugTokenURL() string {
	if s.mode == "debug" {
		return s.vp.GetString("FBLogin.Debug.DebugTokenURL")
	}
	return s.vp.GetString("FBLogin.Release.DebugTokenURL")
}
