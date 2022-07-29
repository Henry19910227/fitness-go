package iap

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

func (s *setting) GetSandboxURL() string {
	return s.vp.GetString("IAP.Sandbox")
}

func (s *setting) GetProductURL() string {
	return s.vp.GetString("IAP.Product")
}

func (s *setting) GetAppServerAPIURL() string {
	if s.mode == "debug" {
		return s.vp.GetString("IAP.Debug.AppStoreServer")
	}
	return s.vp.GetString("IAP.Release.AppStoreServer")
}

func (s *setting) GetPassword() string {
	if s.mode == "debug" {
		return s.vp.GetString("IAP.Debug.Password")
	}
	return s.vp.GetString("IAP.Release.Password")
}

func (s *setting) GetKeyPath() string {
	if s.mode == "debug" {
		return s.vp.GetString("IAP.Debug.KeyPath")
	}
	return s.vp.GetString("IAP.Release.KeyPath")
}

func (s *setting) GetKeyName() string {
	if s.mode == "debug" {
		return s.vp.GetString("IAP.Debug.KeyName")
	}
	return s.vp.GetString("IAP.Release.KeyName")
}

func (s *setting) GetKeyID() string {
	if s.mode == "debug" {
		return s.vp.GetString("IAP.Debug.KeyID")
	}
	return s.vp.GetString("IAP.Release.KeyID")
}

func (s *setting) GetBundleID() string {
	if s.mode == "debug" {
		return s.vp.GetString("IAP.Debug.BundleID")
	}
	return s.vp.GetString("IAP.Release.BundleID")
}

func (s *setting) GetIssuer() string {
	if s.mode == "debug" {
		return s.vp.GetString("IAP.Debug.Issuer")
	}
	return s.vp.GetString("IAP.Release.Issuer")
}
