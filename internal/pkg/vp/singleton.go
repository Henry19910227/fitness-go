package vp

import (
	"github.com/Henry19910227/fitness-go/config"
	"github.com/prometheus/common/log"
	"github.com/spf13/viper"
	"sync"
)

var vp *viper.Viper
var once sync.Once

func Shared() *viper.Viper {
	once.Do(func() {
		vp = viper.New()
		vp.SetConfigFile(config.RootPath() + "/config.yaml")
		if err := vp.ReadInConfig(); err != nil {
			log.Fatalf(err.Error())
		}
	})
	return vp
}
