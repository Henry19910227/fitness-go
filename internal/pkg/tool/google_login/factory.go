package google_login

import setting "github.com/Henry19910227/fitness-go/internal/pkg/setting/google_login"

func NewTool() Tool {
	return New(setting.New())
}
