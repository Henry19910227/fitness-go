package mail

import setting "github.com/Henry19910227/fitness-go/internal/pkg/setting/mail"

func NewTool() Tool {
	return New(setting.New())
}
