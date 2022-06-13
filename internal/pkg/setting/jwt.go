package setting

import (
	"time"

	"github.com/spf13/viper"
)

type jwtSetting struct {
	vp *viper.Viper
}

func NewJWT(viperTool *viper.Viper) JWT {
	return &jwtSetting{viperTool}
}

func (setting *jwtSetting) GetTokenSecret() string {
	return setting.vp.GetString("JWT.Secret")
}

func (setting *jwtSetting) GetIssuer() string {
	return setting.vp.GetString("JWT.Issuer")
}

func (setting *jwtSetting) GetExpire() time.Duration {
	return setting.vp.GetDuration("JWT.Expire") * time.Minute
}