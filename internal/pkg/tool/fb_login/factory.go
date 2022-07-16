package fb_login

import setting "github.com/Henry19910227/fitness-go/internal/pkg/setting/fb_login"

func NewTool() Tool {
	return New(setting.New())
}
