package line_login

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

func (s *setting) GetVerifyTokenURL() string {
	return s.vp.GetString("LineLogin.VerifyTokenURL")
}

func (s *setting) GetProfileURL() string {
	return s.vp.GetString("LineLogin.ProfileURL")
}

func (s *setting) GetClientID() string {
	if s.mode == "debug" {
		return s.vp.GetString("LineLogin.Debug.ClientID")
	}
	if s.mode == "release" {
		return s.vp.GetString("LineLogin.Release.ClientID")
	}
	if s.mode == "production" {
		return s.vp.GetString("LineLogin.Production.ClientID")
	}
	return ""
}
