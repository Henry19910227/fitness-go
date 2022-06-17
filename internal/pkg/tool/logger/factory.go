package logger

import (
	setting "github.com/Henry19910227/fitness-go/internal/pkg/setting/logger"
)

func NewTool() Tool {
	return New(setting.New())
}
