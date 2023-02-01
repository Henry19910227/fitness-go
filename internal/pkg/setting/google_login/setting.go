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

func (s *setting) GetAndroidClientID() string {
	if s.mode == "debug" {
		return s.vp.GetString("GoogleLogin.Debug.Android.ClientID")
	}
	if s.mode == "release" {
		return s.vp.GetString("GoogleLogin.Release.Android.ClientID")
	}
	if s.mode == "production" {
		return s.vp.GetString("GoogleLogin.Production.Android.ClientID")
	}
	return ""
}

func (s *setting) GetIOSClientID() string {
	if s.mode == "debug" {
		return s.vp.GetString("GoogleLogin.Debug.iOS.ClientID")
	}
	if s.mode == "release" {
		return s.vp.GetString("GoogleLogin.Release.iOS.ClientID")
	}
	if s.mode == "production" {
		return s.vp.GetString("GoogleLogin.Production.iOS.ClientID")
	}
	return ""
}

func (s *setting) GetIss() string {
	if s.mode == "debug" {
		return s.vp.GetString("GoogleLogin.Iss")
	}
	if s.mode == "release" {
		return s.vp.GetString("GoogleLogin.Iss")
	}
	if s.mode == "production" {
		return s.vp.GetString("GoogleLogin.Iss")
	}
	return ""
}

func (s *setting) GetDebugTokenURL() string {
	if s.mode == "debug" {
		return s.vp.GetString("GoogleLogin.DebugTokenURL")
	}
	if s.mode == "release" {
		return s.vp.GetString("GoogleLogin.DebugTokenURL")
	}
	if s.mode == "production" {
		return s.vp.GetString("GoogleLogin.DebugTokenURL")
	}
	return ""
}