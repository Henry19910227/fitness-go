package jwt

import setting "github.com/Henry19910227/fitness-go/internal/pkg/setting/jwt"

func NewTool() Tool {
	return New(setting.New())
}
