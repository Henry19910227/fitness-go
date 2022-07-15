package user

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/global"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/orm"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/redis"
	"github.com/Henry19910227/fitness-go/internal/v2/controller/user"
	tokenMiddleware "github.com/Henry19910227/fitness-go/internal/v2/middleware/token"
	"github.com/gin-gonic/gin"
)

func SetRoute(v2 *gin.RouterGroup) {
	controller := user.NewController(orm.Shared().DB())
	midd := tokenMiddleware.NewTokenMiddleware(redis.Shared())
	v2.PATCH("/password", midd.Verify([]global.Role{global.UserRole}), controller.UpdatePassword)
	v2.POST("/login/email", controller.LoginForEmail)
	v2.POST("/logout", midd.Verify([]global.Role{global.UserRole}), controller.Logout)
	v2.POST("/register/email", controller.RegisterForEmail)
	v2.POST("/register/otp", controller.CreateRegisterOTP)
	v2.POST("/register/account/validate", controller.RegisterAccountValidate)
	v2.POST("/register/nickname/validate", controller.RegisterNicknameValidate)
}
