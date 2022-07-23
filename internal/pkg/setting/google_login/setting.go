package google_login

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

func (s *setting) GetClientID() string {
	if s.mode == "debug" {
		return s.vp.GetString("GoogleLogin.Debug.ClientID")
	}
	return s.vp.GetString("GoogleLogin.Release.ClientID")
}

func (s *setting) GetIss() string {
	if s.mode == "debug" {
		return s.vp.GetString("GoogleLogin.Iss")
	}
	return s.vp.GetString("GoogleLogin.Iss")
}

func (s *setting) GetDebugTokenURL() string {
	if s.mode == "debug" {
		return s.vp.GetString("GoogleLogin.Debug.DebugTokenURL")
	}
	return s.vp.GetString("GoogleLogin.Release.DebugTokenURL")
}