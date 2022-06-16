package mysql

import "github.com/spf13/viper"

func NewSetting(vp *viper.Viper) Setting {
	return New(vp)
}

func NewMockSetting() Setting {
	return NewMock()
}
