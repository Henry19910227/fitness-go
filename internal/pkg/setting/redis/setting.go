package redis

import (
	"github.com/Henry19910227/fitness-go/config"
	"github.com/Henry19910227/fitness-go/internal/pkg/build"
	"github.com/prometheus/common/log"
	"github.com/spf13/viper"
)

type setting struct {
	vp   *viper.Viper
	mode string
}

func New() Setting {
	path := config.RootPath() + "/config.yaml"
	vp := viper.New()
	vp.SetConfigFile(path)
	if err := vp.ReadInConfig(); err != nil {
		log.Fatalf(err.Error())
	}
	return &setting{vp: vp, mode: build.RunMode()}
}

func (s *setting) GetHost() string {
	if s.mode == "debug" {
		return s.vp.GetString("Redis.Debug.Host")
	}
	return s.vp.GetString("Redis.Release.Host")
}

func (s *setting) GetPwd() string {
	if s.mode == "debug" {
		return s.vp.GetString("Redis.Debug.Password")
	}
	return s.vp.GetString("Redis.Release.Password")
}
