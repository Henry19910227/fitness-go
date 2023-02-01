package otp

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/build"
	"github.com/Henry19910227/fitness-go/internal/pkg/vp"
	"github.com/spf13/viper"
)

type setting struct {
	vp   *viper.Viper
	mode string
}

func New() Setting {
	return &setting{vp: vp.Shared(), mode: build.RunMode()}
}

func (s *setting) GetPeriod() int {
	if s.mode == "debug" {
		return s.vp.GetInt("OTP.Debug.Period")
	}
	if s.mode == "release" {
		return s.vp.GetInt("OTP.Release.Period")
	}
	if s.mode == "production" {
		return s.vp.GetInt("OTP.Production.Period")
	}
	return 300
}
