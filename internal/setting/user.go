package setting

import (
	"github.com/spf13/viper"
	"time"
)

type user struct {
	vp *viper.Viper
	mode string
}

func NewUser(viperTool *viper.Viper) User {
	return &user{viperTool, viperTool.GetString("Server.RunMode")}
}

func (u *user) GetOnlineExpire() time.Duration  {
	if u.mode == "debug" {
		return u.vp.GetDuration("User.Debug.OnlineExpire") * time.Minute
	}
	return u.vp.GetDuration("User.Release.OnlineExpire") * time.Minute
}
