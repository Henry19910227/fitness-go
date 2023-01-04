package admin

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/jwt"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/redis"
	"github.com/Henry19910227/fitness-go/internal/v2/service/admin"
	"gorm.io/gorm"
)

func NewResolver(db *gorm.DB) Resolver {
	adminService := admin.NewService(db)
	redisTool := redis.Shared()
	jwtTool := jwt.NewTool()
	return New(adminService, redisTool, jwtTool)
}
