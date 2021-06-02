package setting

import "github.com/spf13/viper"

type redis struct {
	vp   *viper.Viper
	mode string
}

func NewRedis(v *viper.Viper) Redis {
	return &redis{v, v.GetString("Server.RunMode")}
}

func (r *redis) GetHost() string {
	if r.mode == "debug" {
		return r.vp.GetString("Redis.Debug.Host")
	}
	return r.vp.GetString("Redis.Release.Host")
}

func (r *redis) GetPwd() string {
	if r.mode == "debug" {
		return r.vp.GetString("Redis.Debug.Password")
	}
	return r.vp.GetString("Redis.Release.Password")
}
