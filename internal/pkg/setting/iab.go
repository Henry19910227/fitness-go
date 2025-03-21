package setting

import "github.com/spf13/viper"

type iab struct {
	vp   *viper.Viper
	mode string
}

func NewIAB(viperTool *viper.Viper) IAB {
	return &iab{viperTool, viperTool.GetString("Server.RunMode")}
}

func (i *iab) GetURL() string {
	return i.vp.GetString("IAB.URL")
}

func (i *iab) GetScope() string {
	return i.vp.GetString("IAB.Scope")
}

func (i *iab) GetJsonFilePath() string {
	if i.mode == "debug" {
		return i.vp.GetString("IAB.Debug.JsonFilePath")
	}
	if i.mode == "release" {
		return i.vp.GetString("IAB.Release.JsonFilePath")
	}
	if i.mode == "production" {
		return i.vp.GetString("IAB.Production.JsonFilePath")
	}
	return ""
}

func (i *iab) GetPackageName() string {
	if i.mode == "debug" {
		return i.vp.GetString("IAB.Debug.PackageName")
	}
	if i.mode == "release" {
		return i.vp.GetString("IAB.Release.PackageName")
	}
	if i.mode == "production" {
		return i.vp.GetString("IAB.Production.PackageName")
	}
	return ""
}
