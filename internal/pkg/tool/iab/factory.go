package iab

import setting "github.com/Henry19910227/fitness-go/internal/pkg/setting/iab"

func NewTool() Tool {
	return New(setting.New())
}
