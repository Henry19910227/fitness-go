package orm

import (
	setting "github.com/Henry19910227/fitness-go/internal/pkg/setting/mysql"
)

func NewTool() Tool {
	return New(setting.New())
}

func NewMockTool() Tool {
	return New(setting.NewMockSetting())
}
