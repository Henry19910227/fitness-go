package otp

import setting "github.com/Henry19910227/fitness-go/internal/pkg/setting/otp"

func NewTool() Tool {
	return New(setting.New())
}
