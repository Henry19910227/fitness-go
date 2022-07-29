package iap

import setting "github.com/Henry19910227/fitness-go/internal/pkg/setting/iap"

func NewTool() Tool {
	return New(setting.New())
}
