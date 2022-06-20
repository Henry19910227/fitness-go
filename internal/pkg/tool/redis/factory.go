package redis

import setting "github.com/Henry19910227/fitness-go/internal/pkg/setting/redis"

func NewTool() Tool {
	return New(setting.New())
}
