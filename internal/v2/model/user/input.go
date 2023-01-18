package user

import (
	pagingOptional "github.com/Henry19910227/fitness-go/internal/v2/field/paging/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/field/user/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/field/user/required"
	"github.com/Henry19910227/fitness-go/internal/v2/model/join"
	orderBy "github.com/Henry19910227/fitness-go/internal/v2/model/order_by"
	"github.com/Henry19910227/fitness-go/internal/v2/model/preload"
	"github.com/Henry19910227/fitness-go/internal/v2/model/where"
	"mime/multipart"
)

type GenerateInput struct {
	DataAmount int
}

type PagingInput = struct {
	pagingOptional.PageField
	pagingOptional.SizeField
}
type PreloadInput = preload.Input
type WhereInput = where.Input
type JoinInput = join.Input
type OrderByInput = orderBy.Input
type CustomOrderByInput = orderBy.CustomInput

type ListInput struct {
	optional.IDField
	optional.AccountField
	optional.PasswordField
	optional.UserStatusField
	optional.UserTypeField
	optional.NicknameField
	optional.EmailField
	optional.IsDeletedField
	JoinInput
	WhereInput
	PagingInput
	PreloadInput
	OrderByInput
	CustomOrderByInput
}

type FindInput struct {
	optional.IDField
	optional.IsDeletedField
	PreloadInput
}

// APIUpdatePasswordInput /v2/password [PATCH]
type APIUpdatePasswordInput struct {
	required.IDField
	Body APIUpdatePasswordBody
}
type APIUpdatePasswordBody struct {
	OldPassword string `json:"old_password" binding:"required,min=6,max=18" example:"12345678"` // 舊密碼 (6~18字元)
	required.PasswordField
}

// APIUpdateUserProfileInput /v2/user/profile [PATCH]
type APIUpdateUserProfileInput struct {
	required.IDField
	Body APIUpdateUserProfileBody
}
type APIUpdateUserProfileBody struct {
	optional.DeviceTokenField
	optional.NicknameField
	optional.SexField
	optional.HeightField
	optional.WeightField
	optional.BirthdayField
	optional.ExperienceField
	optional.TargetField
}

// APIUpdateUserAvatarInput /v2/user/avatar [PATCH]
type APIUpdateUserAvatarInput struct {
	required.IDField
	CoverNamed string
	File       multipart.File
}

// APIGetUserProfileInput /v2/user/profile [GET]
type APIGetUserProfileInput struct {
	required.IDField
}

// APILoginForEmailInput /v2/login/email [POST]
type APILoginForEmailInput struct {
	Body APILoginForEmailBody
}
type APILoginForEmailBody struct {
	required.EmailField
	required.PasswordField
}

// APIGetAppleRefreshTokenInput /v2/apple_refresh_token [POST]
type APIGetAppleRefreshTokenInput struct {
	Body APIGetAppleRefreshTokenBody
}
type APIGetAppleRefreshTokenBody struct {
	AccessToken string `json:"access_token" binding:"required" example:"EAAucgU8qZCzMBAOZCy59TLD1aM2NAO1ITBpZC64imFp95CRuPv4ZAWepAMV"` // apple sdk 回傳的 authorizationCode
}

// APILoginForFacebookInput /v2/login/facebook [POST]
type APILoginForFacebookInput struct {
	Body APILoginForFacebookBody
}
type APILoginForFacebookBody struct {
	AccessToken string `json:"access_token" binding:"required" example:"EAAucgU8qZCzMBAOZCy59TLD1aM2NAO1ITBpZC64imFp95CRuPv4ZAWepAMV"` // sdk 回傳的 token string
}

// APILoginForGoogleInput /v2/login/google [POST]
type APILoginForGoogleInput struct {
	Body APILoginForGoogleBody
}
type APILoginForGoogleBody struct {
	AccessToken string `json:"access_token" binding:"required" example:"EAAucgU8qZCzMBAOZCy59TLD1aM2NAO1ITBpZC64imFp95CRuPv4ZAWepAMV"` // sdk 回傳的 token string
}

// APILoginForAppleInput /v2/login/apple [POST]
type APILoginForAppleInput struct {
	Body APILoginForAppleBody
}
type APILoginForAppleBody struct {
	RefreshToken string `json:"refresh_token" binding:"required" example:"EAAucgU8qZCzMBAOZCy59TLD1aM2NAO1ITBpZC64imFp95CRuPv4ZAWepAMV"` // refresh token
}

// APILoginForLineInput /v2/login/line [POST]
type APILoginForLineInput struct {
	Body APILoginForLineBody
}
type APILoginForLineBody struct {
	AccessToken string `json:"access_token" binding:"required" example:"EAAucgU8qZCzMBAOZCy59TLD1aM2NAO1ITBpZC64imFp95CRuPv4ZAWepAMV"` // sdk 回傳的 token string
}

// APILogoutInput /v2/logout [POST]
type APILogoutInput struct {
	required.IDField
}

// APIRegisterForEmailInput /v2/register/email [POST]
type APIRegisterForEmailInput struct {
	Body APIRegisterForEmailBody
}
type APIRegisterForEmailBody struct {
	required.EmailField
	required.PasswordField
	required.NicknameField
	OTPCode string `json:"otp_code" binding:"required,max=16" example:"531476"` // 信箱驗證碼
}

// APIRegisterForFacebookInput /v2/register/facebook [POST]
type APIRegisterForFacebookInput struct {
	Body APIRegisterForFacebookBody
}
type APIRegisterForFacebookBody struct {
	AccessToken string `json:"access_token" binding:"required" example:"EAAucgU8qZCzMBAOZCy59TLD1aM2NAO1ITBpZC64imFp95CRuPv4ZAWepAMV"` // sdk 回傳的 token string
	required.NicknameField
	required.EmailField
	OTPCode string `json:"otp_code" binding:"required,max=16" example:"531476"` // 信箱驗證碼
}

// APIRegisterForGoogleInput /v2/register/google [POST]
type APIRegisterForGoogleInput struct {
	Body APIRegisterForGoogleBody
}
type APIRegisterForGoogleBody struct {
	AccessToken string `json:"access_token" binding:"required" example:"EAAucgU8qZCzMBAOZCy59TLD1aM2NAO1ITBpZC64imFp95CRuPv4ZAWepAMV"` // sdk 回傳的 token string
	required.NicknameField
	required.EmailField
	OTPCode string `json:"otp_code" binding:"required,max=16" example:"531476"` // 信箱驗證碼
}

// APIRegisterForAppleInput /v2/register/apple [POST]
type APIRegisterForAppleInput struct {
	Body APIRegisterForAppleBody
}
type APIRegisterForAppleBody struct {
	RefreshToken string `json:"refresh_token" binding:"required" example:"EAAucgU8qZCzMBAOZCy59TLD1aM2NAO1ITBpZC64imFp95CRuPv4ZAWepAMV"` // refresh token
	required.NicknameField
	required.EmailField
	OTPCode string `json:"otp_code" binding:"required,max=16" example:"531476"` // 信箱驗證碼
}

// APIRegisterForLineInput /v2/register/line [POST]
type APIRegisterForLineInput struct {
	Body APIRegisterForLineBody
}
type APIRegisterForLineBody struct {
	AccessToken string `json:"access_token" binding:"required" example:"EAAucgU8qZCzMBAOZCy59TLD1aM2NAO1ITBpZC64imFp95CRuPv4ZAWepAMV"` // sdk 回傳的 token string
	required.NicknameField
	required.EmailField
	OTPCode string `json:"otp_code" binding:"required,max=16" example:"531476"` // 信箱驗證碼
}

// APICreateOTPInput /v2/otp [POST]
type APICreateOTPInput struct {
	Body APICreateOTPBody
}
type APICreateOTPBody struct {
	required.EmailField
}

// APIRegisterNicknameValidateInput /v2/register/nickname/validate [POST]
type APIRegisterNicknameValidateInput struct {
	Body APIRegisterNicknameValidateBody
}
type APIRegisterNicknameValidateBody struct {
	required.NicknameField
}

// APIRegisterEmailValidateInput /v2/register/email/validate [POST]
type APIRegisterEmailValidateInput struct {
	Body APIRegisterEmailValidateBody
}
type APIRegisterEmailValidateBody struct {
	required.EmailField
}

// APIRegisterEmailAccountValidateInput /v2/register/email_account/validate [POST]
type APIRegisterEmailAccountValidateInput struct {
	Body APIRegisterEmailAccountValidateBody
}
type APIRegisterEmailAccountValidateBody struct {
	required.EmailField
}

// APIRegisterFacebookAccountValidateInput /v2/register/facebook_account/validate [POST]
type APIRegisterFacebookAccountValidateInput struct {
	Body APIRegisterFacebookAccountValidateBody
}
type APIRegisterFacebookAccountValidateBody struct {
	AccessToken string `json:"access_token" binding:"required" example:"EAAucgU8qZCzMBAOZCy59TLD1aM2NAO1ITBpZC64imFp95CRuPv4ZAWepAMV"` // sdk 回傳的 token string
}

// APIRegisterLineAccountValidateInput /v2/register/line_account/validate [POST]
type APIRegisterLineAccountValidateInput struct {
	Body APIRegisterLineAccountValidateBody
}
type APIRegisterLineAccountValidateBody struct {
	AccessToken string `json:"access_token" binding:"required" example:"EAAucgU8qZCzMBAOZCy59TLD1aM2NAO1ITBpZC64imFp95CRuPv4ZAWepAMV"` // sdk 回傳的 token string
}

// APIRegisterGoogleAccountValidateInput /v2/register/google_account/validate [POST]
type APIRegisterGoogleAccountValidateInput struct {
	Body APIRegisterGoogleAccountValidateBody
}
type APIRegisterGoogleAccountValidateBody struct {
	AccessToken string `json:"access_token" binding:"required" example:"EAAucgU8qZCzMBAOZCy59TLD1aM2NAO1ITBpZC64imFp95CRuPv4ZAWepAMV"` // sdk 回傳的 token string
}

// APIRegisterAppleAccountValidateInput /v2/register/apple_account/validate [POST]
type APIRegisterAppleAccountValidateInput struct {
	Body APIRegisterAppleAccountValidateBody
}
type APIRegisterAppleAccountValidateBody struct {
	RefreshToken string `json:"refresh_token" binding:"required" example:"EAAucgU8qZCzMBAOZCy59TLD1aM2NAO1ITBpZC64imFp95CRuPv4ZAWepAMV"` // 透過 authorizationCode 取得的 refresh token
}

// APICreateResetOTPInput /v2/reset_password/otp [POST]
type APICreateResetOTPInput struct {
	Body APICreateResetOTPBody
}
type APICreateResetOTPBody struct {
	Email string `json:"email" binding:"required,email" example:"henry@gmail.com"` // 信箱
}

// APIResetOTPValidateInput /v2/reset_password/otp_validate [POST]
type APIResetOTPValidateInput struct {
	Body APIResetOTPValidateBody
}
type APIResetOTPValidateBody struct {
	OTPCode string `json:"otp_code" binding:"required,max=16" example:"531476"`      // 信箱驗證碼
	Email   string `json:"email" binding:"required,email" example:"henry@gmail.com"` // 信箱
}

// APIUpdateResetPasswordInput /v2/reset_password/password [PATCH]
type APIUpdateResetPasswordInput struct {
	Body APIUpdateResetPasswordBody
}
type APIUpdateResetPasswordBody struct {
	OTPCode string `json:"otp_code" binding:"required,max=16" example:"531476"`      // 信箱驗證碼
	Email   string `json:"email" binding:"required,email" example:"henry@gmail.com"` // 信箱
	required.PasswordField
}

// APIDeleteUserInput /v2/user [DELETE]
type APIDeleteUserInput struct {
	required.UserIDField
}
