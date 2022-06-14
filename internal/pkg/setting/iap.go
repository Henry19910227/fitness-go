package setting

import (
	"github.com/spf13/viper"
)

type iap struct {
	vp   *viper.Viper
	mode string
}

func NewIAP(viperTool *viper.Viper) IAP {
	return &iap{viperTool, viperTool.GetString("Server.RunMode")}
}

func (i *iap) GetSandboxURL() string {
	return i.vp.GetString("IAP.Sandbox")
}

func (i *iap) GetProductURL() string {
	return i.vp.GetString("IAP.Product")
}

func (i *iap) GetAppServerAPIURL() string {
	if i.mode == "debug" {
		return i.vp.GetString("IAP.Debug.AppStoreServer")
	}
	return i.vp.GetString("IAP.Release.AppStoreServer")
}

func (i *iap) GetPassword() string {
	if i.mode == "debug" {
		return i.vp.GetString("IAP.Debug.Password")
	}
	return i.vp.GetString("IAP.Release.Password")
}

func (i *iap) GetKeyPath() string {
	if i.mode == "debug" {
		return i.vp.GetString("IAP.Debug.KeyPath")
	}
	return i.vp.GetString("IAP.Release.KeyPath")
}

func (i *iap) GetKeyID() string {
	if i.mode == "debug" {
		return i.vp.GetString("IAP.Debug.KeyID")
	}
	return i.vp.GetString("IAP.Release.KeyID")
}

func (i *iap) GetBundleID() string {
	if i.mode == "debug" {
		return i.vp.GetString("IAP.Debug.BundleID")
	}
	return i.vp.GetString("IAP.Release.BundleID")
}

func (i *iap) GetIssuer() string {
	if i.mode == "debug" {
		return i.vp.GetString("IAP.Debug.Issuer")
	}
	return i.vp.GetString("IAP.Release.Issuer")
}
