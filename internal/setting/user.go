package setting

import (
	"github.com/spf13/viper"
)

type user struct {
	vp *viper.Viper
	mode string
}

func NewUser(viperTool *viper.Viper) User {
	return &user{viperTool, viperTool.GetString("Server.RunMode")}
}