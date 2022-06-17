package jwt

import (
	"github.com/Henry19910227/fitness-go/config"
	"github.com/prometheus/common/log"
	"github.com/spf13/viper"
	"time"
)

type setting struct {
	vp *viper.Viper
}

func New() Setting {
	vp := viper.New()
	vp.SetConfigFile(config.RootPath() + "/config.yaml")
	if err := vp.ReadInConfig(); err != nil {
		log.Fatalf(err.Error())
	}
	return &setting{vp}
}

func (s *setting) GetTokenSecret() string {
	return s.vp.GetString("JWT.Secret")
}

func (s *setting) GetIssuer() string {
	return s.vp.GetString("JWT.Issuer")
}

func (s *setting) GetExpire() time.Duration {
	return s.vp.GetDuration("JWT.Expire") * time.Minute
}
