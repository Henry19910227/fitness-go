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

func (i *iap) GetPassword() string {
	if i.mode == "debug" {
		return i.vp.GetString("IAP.Debug.Password")
	}
	return i.vp.GetString("IAP.Release.Password")
}
