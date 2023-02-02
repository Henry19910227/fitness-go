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
	if i.mode == "release" {
		return i.vp.GetString("IAP.Release.AppStoreServer")
	}
	if i.mode == "production" {
		return i.vp.GetString("IAP.Production.AppStoreServer")
	}
	return ""
}

func (i *iap) GetPassword() string {
	if i.mode == "debug" {
		return i.vp.GetString("IAP.Debug.Password")
	}
	if i.mode == "release" {
		return i.vp.GetString("IAP.Release.Password")
	}
	if i.mode == "production" {
		return i.vp.GetString("IAP.Production.Password")
	}
	return ""
}

func (i *iap) GetKeyPath() string {
	if i.mode == "debug" {
		return i.vp.GetString("IAP.Debug.KeyPath")
	}
	if i.mode == "release" {
		return i.vp.GetString("IAP.Release.KeyPath")
	}
	if i.mode == "production" {
		return i.vp.GetString("IAP.Production.KeyPath")
	}
	return ""
}

func (i *iap) GetKeyID() string {
	if i.mode == "debug" {
		return i.vp.GetString("IAP.Debug.KeyID")
	}
	if i.mode == "release" {
		return i.vp.GetString("IAP.Release.KeyID")
	}
	if i.mode == "production" {
		return i.vp.GetString("IAP.Production.KeyID")
	}
	return ""
}

func (i *iap) GetBundleID() string {
	if i.mode == "debug" {
		return i.vp.GetString("IAP.Debug.BundleID")
	}
	if i.mode == "release" {
		return i.vp.GetString("IAP.Release.BundleID")
	}
	if i.mode == "production" {
		return i.vp.GetString("IAP.Production.BundleID")
	}
	return ""
}

func (i *iap) GetIssuer() string {
	if i.mode == "debug" {
		return i.vp.GetString("IAP.Debug.Issuer")
	}
	if i.mode == "release" {
		return i.vp.GetString("IAP.Release.Issuer")
	}
	if i.mode == "production" {
		return i.vp.GetString("IAP.Production.Issuer")
	}
	return ""
}
