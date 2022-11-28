package fcm

import setting "github.com/Henry19910227/fitness-go/internal/pkg/setting/fcm"

func NewTool() Tool {
	return New(setting.New())
}
