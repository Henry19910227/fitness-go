package user

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/user"
	"github.com/Henry19910227/fitness-go/internal/v2/model/user/api_get_cms_course_users"
	"github.com/Henry19910227/fitness-go/internal/v2/model/user/api_get_cms_user"
	"github.com/Henry19910227/fitness-go/internal/v2/model/user/api_get_cms_users"
	"github.com/Henry19910227/fitness-go/internal/v2/model/user/api_update_cms_user"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Resolver interface {
	APIGetCMSCourseUsers(input *api_get_cms_course_users.Input) (output api_get_cms_course_users.Output)
	APIGetCMSUser(input *api_get_cms_user.Input) (output api_get_cms_user.Output)
	APIGetCMSUsers(input *api_get_cms_users.Input) (output api_get_cms_users.Output)
	APIUpdateCMSUser(input *api_update_cms_user.Input) (output api_update_cms_user.Output)

	APIUpdatePassword(input *model.APIUpdatePasswordInput) (output model.APIUpdatePasswordOutput)
	APIUpdateUserProfile(input *model.APIUpdateUserProfileInput) (output model.APIUpdateUserProfileOutput)
	APIUpdateUserAvatar(input *model.APIUpdateUserAvatarInput) (output model.APIUpdateUserAvatarOutput)
	APIGetUserProfile(input *model.APIGetUserProfileInput) (output model.APIGetUserProfileOutput)
	APIGetAppleRefreshToken(input *model.APIGetAppleRefreshTokenInput) (output model.APIGetAppleRefreshTokenOutput)
	APIRegisterForEmail(input *model.APIRegisterForEmailInput) (output model.APIRegisterForEmailOutput)
	APIRegisterForFacebook(input *model.APIRegisterForFacebookInput) (output model.APIRegisterForFacebookOutput)
	APIRegisterForGoogle(input *model.APIRegisterForGoogleInput) (output model.APIRegisterForGoogleOutput)
	APIRegisterForApple(input *model.APIRegisterForAppleInput) (output model.APIRegisterForAppleOutput)
	APIRegisterForLine(input *model.APIRegisterForLineInput) (output model.APIRegisterForLineOutput)
	APILoginForEmail(input *model.APILoginForEmailInput) (output model.APILoginForEmailOutput)
	APILoginForFacebook(input *model.APILoginForFacebookInput) (output model.APILoginForFacebookOutput)
	APILoginForGoogle(input *model.APILoginForGoogleInput) (output model.APILoginForGoogleOutput)
	APILoginForLine(input *model.APILoginForLineInput) (output model.APILoginForLineOutput)
	APILoginForApple(input *model.APILoginForAppleInput) (output model.APILoginForAppleOutput)
	APILogout(input *model.APILogoutInput) (output model.APILogoutOutput)
	APICreateRegisterOTP(input *model.APICreateOTPInput) (output model.APICreateRegisterOTPOutput)
	APIRegisterNicknameValidate(input *model.APIRegisterNicknameValidateInput) (output model.APIRegisterNicknameValidateOutput)
	APIRegisterEmailValidate(input *model.APIRegisterEmailValidateInput) (output model.APIRegisterEmailValidateOutput)
	APIRegisterEmailAccountValidate(input *model.APIRegisterEmailAccountValidateInput) (output model.APIRegisterEmailAccountValidateOutput)
	APIRegisterFacebookAccountValidate(input *model.APIRegisterFacebookAccountValidateInput) (output model.APIRegisterFacebookAccountValidateOutput)
	APIRegisterLineAccountValidate(input *model.APIRegisterLineAccountValidateInput) (output model.APIRegisterLineAccountValidateOutput)
	APIRegisterGoogleAccountValidate(input *model.APIRegisterGoogleAccountValidateInput) (output model.APIRegisterGoogleAccountValidateOutput)
	APIRegisterAppleAccountValidate(input *model.APIRegisterAppleAccountValidateInput) (output model.APIRegisterAppleAccountValidateOutput)
	APICreateResetOTP(input *model.APICreateResetOTPInput) (output model.APICreateResetOTPOutput)
	APIResetOTPValidate(input *model.APIResetOTPValidateInput) (output model.APIResetOTPValidateOutput)
	APIUpdateResetPassword(input *model.APIUpdateResetPasswordInput) (output model.APIUpdateResetPasswordOutput)
	APIDeleteUser(ctx *gin.Context, tx *gorm.DB, input *model.APIDeleteUserInput) (output model.APIDeleteUserOutput)
}
