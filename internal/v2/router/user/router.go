package user

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/global"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/orm"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/redis"
	"github.com/Henry19910227/fitness-go/internal/v2/controller/user"
	"github.com/Henry19910227/fitness-go/internal/v2/middleware"
	tokenMiddleware "github.com/Henry19910227/fitness-go/internal/v2/middleware/token"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetRoute(v2 *gin.RouterGroup) {
	controller := user.NewController(orm.Shared().DB())
	midd := tokenMiddleware.NewTokenMiddleware(redis.Shared())
	v2.StaticFS("/resource/user/avatar", http.Dir("./volumes/storage/user/avatar"))
	v2.GET("/cms/course/:course_id/users", midd.Verify([]global.Role{global.AdminRole}), controller.GetCMSCourseUsers)
	v2.GET("/cms/user/:user_id", midd.Verify([]global.Role{global.AdminRole}), controller.GetCMSUser)

	v2.PATCH("/password", midd.Verify([]global.Role{global.UserRole}), controller.UpdatePassword)
	v2.PATCH("/user/profile", midd.Verify([]global.Role{global.UserRole}), controller.UpdateUserProfile)
	v2.PATCH("/user/avatar", midd.Verify([]global.Role{global.UserRole}), controller.UpdateUserAvatar)
	v2.GET("/user/profile", midd.Verify([]global.Role{global.UserRole}), controller.GetUserProfile)
	v2.POST("/login/email", controller.LoginForEmail)
	v2.POST("/login/facebook", controller.LoginForFacebook)
	v2.POST("/login/google", controller.LoginForGoogle)
	v2.POST("/login/line", controller.LoginForLine)
	v2.POST("/login/apple", controller.LoginForApple)
	v2.POST("/logout", midd.Verify([]global.Role{global.UserRole}), controller.Logout)
	v2.POST("/apple_refresh_token", controller.GetAppleRefreshToken)
	v2.POST("/register/email", controller.RegisterForEmail)
	v2.POST("/register/facebook", controller.RegisterForFacebook)
	v2.POST("/register/google", controller.RegisterForGoogle)
	v2.POST("/register/line", controller.RegisterForLine)
	v2.POST("/register/apple", controller.RegisterForApple)
	v2.POST("/otp", controller.CreateRegisterOTP)
	v2.POST("/register/email_account/validate", controller.RegisterEmailAccountValidate)
	v2.POST("/register/facebook_account/validate", controller.RegisterFacebookAccountValidate)
	v2.POST("/register/line_account/validate", controller.RegisterLineAccountValidate)
	v2.POST("/register/google_account/validate", controller.RegisterGoogleAccountValidate)
	v2.POST("/register/apple_account/validate", controller.RegisterAppleAccountValidate)
	v2.POST("/register/nickname/validate", controller.RegisterNicknameValidate)
	v2.POST("/register/email/validate", controller.RegisterEmailValidate)
	v2.POST("/reset_password/otp", controller.CreateResetOTP)
	v2.POST("/reset_password/otp_validate", controller.ResetOTPValidate)
	v2.PATCH("/reset_password/password", controller.UpdateResetPassword)
	v2.DELETE("/user", middleware.Transaction(orm.Shared().DB()), midd.Verify([]global.Role{global.UserRole}), controller.DeleteUser)

}
