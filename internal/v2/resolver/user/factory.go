package user

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/otp"
	userService "github.com/Henry19910227/fitness-go/internal/v2/service/user"
	"gorm.io/gorm"
)

func NewResolver(db *gorm.DB) Resolver {
	userSvc := userService.NewService(db)
	otpTool := otp.New()
	return New(userSvc, otpTool)
}
