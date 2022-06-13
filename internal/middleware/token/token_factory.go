package token

import (
	"github.com/Henry19910227/fitness-go/internal/setting"
	"github.com/Henry19910227/fitness-go/internal/tool"
	"github.com/spf13/viper"
)

func NewTokenMiddleware(redisTool tool.Redis, viperTool *viper.Viper) Middleware {
	jwtTool := tool.NewJWT(setting.NewJWT(viperTool))
	return NewMiddleware(jwtTool, redisTool)
}
