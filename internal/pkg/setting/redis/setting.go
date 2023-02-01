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
	if s.mode == "release" {
		return s.vp.GetString("Redis.Release.Host")
	}
	if s.mode == "production" {
		return s.vp.GetString("Redis.Production.Host")
	}
	return ""
}

func (s *setting) GetPwd() string {
	if s.mode == "debug" {
		return s.vp.GetString("Redis.Debug.Password")
	}
	if s.mode == "release" {
		return s.vp.GetString("Redis.Release.Password")
	}
	if s.mode == "production" {
		return s.vp.GetString("Redis.Production.Password")
	}
	return ""
}
