package user

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/apple_login"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/crypto"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/fb_login"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/google_login"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/iab"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/iap"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/jwt"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/line_login"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/mail"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/otp"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/redis"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/uploader"
	"github.com/Henry19910227/fitness-go/internal/v2/service/course"
	"github.com/Henry19910227/fitness-go/internal/v2/service/receipt"
	"github.com/Henry19910227/fitness-go/internal/v2/service/user"
	"github.com/Henry19910227/fitness-go/internal/v2/service/user_subscribe_info"
	"gorm.io/gorm"
)

func NewResolver(db *gorm.DB) Resolver {
	userService := user.NewService(db)
	receiptService := receipt.NewService(db)
	subscribeInfoService := user_subscribe_info.NewService(db)
	courseService := course.NewService(db)
	otpTool := otp.NewTool()
	cryptoTool := crypto.New()
	redisTool := redis.Shared()
	jwtTool := jwt.NewTool()
	fbLoginTool := fb_login.NewTool()
	googleLoginTool := google_login.NewTool()
	lineLoginTool := line_login.NewTool()
	appleLoginTool := apple_login.NewTool()
	uploadTool := uploader.NewUserAvatarTool()
	iapTool := iap.NewTool()
	iabTool := iab.NewTool()
	mailTool := mail.NewTool()
	return New(userService, receiptService, subscribeInfoService, courseService,
		otpTool, cryptoTool, redisTool, jwtTool,
		fbLoginTool, googleLoginTool, appleLoginTool, lineLoginTool,
		uploadTool, iapTool, iabTool, mailTool)
}
