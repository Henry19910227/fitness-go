package token

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/tool"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/jwt"
)

func NewTokenMiddleware(redisTool tool.Redis) Middleware {
	return NewMiddleware(jwt.NewTool(), redisTool)
}
