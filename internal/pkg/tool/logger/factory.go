package logger

import (
	setting "github.com/Henry19910227/fitness-go/internal/pkg/setting/logger"
	"github.com/spf13/viper"
)

func NewTool(vp *viper.Viper) Tool {
	return New(setting.New(vp))
}
