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
	if s.mode == "release" {
		return s.vp.GetString("IAP.Release.AppStoreServer")
	}
	if s.mode == "production" {
		return s.vp.GetString("IAP.Production.AppStoreServer")
	}
	return ""
}

func (s *setting) GetPassword() string {
	if s.mode == "debug" {
		return s.vp.GetString("IAP.Debug.Password")
	}
	if s.mode == "release" {
		return s.vp.GetString("IAP.Release.Password")
	}
	if s.mode == "production" {
		return s.vp.GetString("IAP.Production.Password")
	}
	return ""
}

func (s *setting) GetKeyPath() string {
	if s.mode == "debug" {
		return s.vp.GetString("IAP.Debug.KeyPath")
	}
	if s.mode == "release" {
		return s.vp.GetString("IAP.Release.KeyPath")
	}
	if s.mode == "production" {
		return s.vp.GetString("IAP.Production.KeyPath")
	}
	return ""
}

func (s *setting) GetKeyName() string {
	if s.mode == "debug" {
		return s.vp.GetString("IAP.Debug.KeyName")
	}
	if s.mode == "release" {
		return s.vp.GetString("IAP.Release.KeyName")
	}
	if s.mode == "production" {
		return s.vp.GetString("IAP.Production.KeyName")
	}
	return ""
}

func (s *setting) GetKeyID() string {
	if s.mode == "debug" {
		return s.vp.GetString("IAP.Debug.KeyID")
	}
	if s.mode == "release" {
		return s.vp.GetString("IAP.Release.KeyID")
	}
	if s.mode == "production" {
		return s.vp.GetString("IAP.Production.KeyID")
	}
	return ""
}

func (s *setting) GetBundleID() string {
	if s.mode == "debug" {
		return s.vp.GetString("IAP.Debug.BundleID")
	}
	if s.mode == "release" {
		return s.vp.GetString("IAP.Release.BundleID")
	}
	if s.mode == "production" {
		return s.vp.GetString("IAP.Production.BundleID")
	}
	return ""
}

func (s *setting) GetIssuer() string {
	if s.mode == "debug" {
		return s.vp.GetString("IAP.Debug.Issuer")
	}
	if s.mode == "release" {
		return s.vp.GetString("IAP.Release.Issuer")
	}
	if s.mode == "production" {
		return s.vp.GetString("IAP.Production.Issuer")
	}
	return ""
}
