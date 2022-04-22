package setting

import "github.com/spf13/viper"

type iab struct {
	vp   *viper.Viper
	mode string
}

func NewIAB(viperTool *viper.Viper) IAB {
	return &iab{viperTool, viperTool.GetString("Server.RunMode")}
}

func (i *iab) GetJsonFilePath() string {
	if i.mode == "debug" {
		return i.vp.GetString("IAB.Debug.JsonFilePath")
	}
	return i.vp.GetString("IAB.Release.JsonFilePath")
}
