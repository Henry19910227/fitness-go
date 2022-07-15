package user

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/crypto"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/jwt"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/otp"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/redis"
	userService "github.com/Henry19910227/fitness-go/internal/v2/service/user"
	"gorm.io/gorm"
)

func NewResolver(db *gorm.DB) Resolver {
	userSvc := userService.NewService(db)
	otpTool := otp.New()
	cryptoTool := crypto.New()
	jwtTool := jwt.NewTool()
	return New(userSvc, otpTool, cryptoTool, redis.Shared(), jwtTool)
}
