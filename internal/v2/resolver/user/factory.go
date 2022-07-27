package user

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/crypto"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/fb_login"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/google_login"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/jwt"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/line_login"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/otp"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/redis"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/uploader"
	userService "github.com/Henry19910227/fitness-go/internal/v2/service/user"
	"gorm.io/gorm"
)

func NewResolver(db *gorm.DB) Resolver {
	userSvc := userService.NewService(db)
	otpTool := otp.New()
	cryptoTool := crypto.New()
	redisTool := redis.Shared()
	jwtTool := jwt.NewTool()
	fbLoginTool := fb_login.NewTool()
	googleLoginTool := google_login.NewTool()
	lineLoginTool := line_login.NewTool()
	uploadTool := uploader.NewUserAvatarTool()
	return New(userSvc, otpTool, cryptoTool, redisTool, jwtTool, fbLoginTool, googleLoginTool, lineLoginTool, uploadTool)
}
