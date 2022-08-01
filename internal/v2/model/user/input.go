package user

import (
	orderBy "github.com/Henry19910227/fitness-go/internal/v2/model/order_by"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"github.com/Henry19910227/fitness-go/internal/v2/model/preload"
	"mime/multipart"
)

type GenerateInput struct {
	DataAmount int
}

type PagingInput = paging.Input
type PreloadInput = preload.Input
type OrderByInput = orderBy.Input

type ListInput struct {
	AccountOptional
	PasswordOptional
	NicknameOptional
	EmailOptional
	IsDeletedOptional
	PagingInput
	PreloadInput
	OrderByInput
}

type FindInput struct {
	IDOptional
	IsDeletedOptional
}

// APIUpdatePasswordInput /v2/password [PATCH]
type APIUpdatePasswordInput struct {
	IDRequired
	Body APIUpdatePasswordBody
}
type APIUpdatePasswordBody struct {
	OldPassword string `json:"old_password" binding:"required,min=6,max=18" example:"12345678"` // 舊密碼 (6~18字元)
	PasswordRequired
}

// APIUpdateUserProfileInput /v2/user/profile [PATCH]
type APIUpdateUserProfileInput struct {
	IDRequired
	Body APIUpdateUserProfileBody
}
type APIUpdateUserProfileBody struct {
	NicknameOptional
	SexOptional
	HeightOptional
	WeightOptional
	BirthdayOptional
	ExperienceOptional
	TargetOptional
}

// APIUpdateUserAvatarInput /v2/user/avatar [PATCH]
type APIUpdateUserAvatarInput struct {
	IDRequired
	CoverNamed string
	File       multipart.File
}

// APIGetUserProfileInput /v2/user/profile [GET]
type APIGetUserProfileInput struct {
	IDRequired
}

// APILoginForEmailInput /v2/login/email [POST]
type APILoginForEmailInput struct {
	Body APILoginForEmailBody
}
type APILoginForEmailBody struct {
	EmailRequired
	PasswordRequired
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
	AccessToken string `json:"access_token" binding:"required" example:"EAAucgU8qZCzMBAOZCy59TLD1aM2NAO1ITBpZC64imFp95CRuPv4ZAWepAMV"` // sdk 回傳的 authorizationCode string
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
	IDRequired
}

// APIRegisterForEmailInput /v2/register/email [POST]
type APIRegisterForEmailInput struct {
	Body APIRegisterForEmailBody
}
type APIRegisterForEmailBody struct {
	EmailRequired
	PasswordRequired
	NicknameRequired
	OTPCode string `json:"otp_code" binding:"required,max=16" example:"531476"` // 信箱驗證碼
}

// APIRegisterForFacebookInput /v2/register/facebook [POST]
type APIRegisterForFacebookInput struct {
	Body APIRegisterForFacebookBody
}
type APIRegisterForFacebookBody struct {
	AccessToken string `json:"access_token" binding:"required" example:"EAAucgU8qZCzMBAOZCy59TLD1aM2NAO1ITBpZC64imFp95CRuPv4ZAWepAMV"` // sdk 回傳的 token string
	NicknameRequired
	EmailRequired
}

// APIRegisterForGoogleInput /v2/register/google [POST]
type APIRegisterForGoogleInput struct {
	Body APIRegisterForGoogleBody
}
type APIRegisterForGoogleBody struct {
	AccessToken string `json:"access_token" binding:"required" example:"EAAucgU8qZCzMBAOZCy59TLD1aM2NAO1ITBpZC64imFp95CRuPv4ZAWepAMV"` // sdk 回傳的 token string
	NicknameRequired
	EmailRequired
}

// APIRegisterForAppleInput /v2/register/apple [POST]
type APIRegisterForAppleInput struct {
	Body APIRegisterForAppleBody
}
type APIRegisterForAppleBody struct {
	AccessToken string `json:"access_token" binding:"required" example:"EAAucgU8qZCzMBAOZCy59TLD1aM2NAO1ITBpZC64imFp95CRuPv4ZAWepAMV"` // sdk 回傳的 authorizationCode string
	NicknameRequired
	EmailRequired
}

// APIRegisterForLineInput /v2/register/line [POST]
type APIRegisterForLineInput struct {
	Body APIRegisterForLineBody
}
type APIRegisterForLineBody struct {
	AccessToken string `json:"access_token" binding:"required" example:"EAAucgU8qZCzMBAOZCy59TLD1aM2NAO1ITBpZC64imFp95CRuPv4ZAWepAMV"` // sdk 回傳的 token string
	NicknameRequired
	EmailRequired
}

// APICreateOTPInput /v2/otp [POST]
type APICreateOTPInput struct {
	Body APICreateOTPBody
}
type APICreateOTPBody struct {
	EmailRequired
}

// APIRegisterNicknameValidateInput /v2/register/nickname/validate [POST]
type APIRegisterNicknameValidateInput struct {
	Body APIRegisterNicknameValidateBody
}
type APIRegisterNicknameValidateBody struct {
	NicknameRequired
}

// APIRegisterEmailValidateInput /v2/register/email/validate [POST]
type APIRegisterEmailValidateInput struct {
	Body APIRegisterEmailValidateBody
}
type APIRegisterEmailValidateBody struct {
	EmailRequired
}

// APIRegisterEmailAccountValidateInput /v2/register/email_account/validate [POST]
type APIRegisterEmailAccountValidateInput struct {
	Body APIRegisterEmailAccountValidateBody
}
type APIRegisterEmailAccountValidateBody struct {
	EmailRequired
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
	AccessToken string `json:"access_token" binding:"required" example:"EAAucgU8qZCzMBAOZCy59TLD1aM2NAO1ITBpZC64imFp95CRuPv4ZAWepAMV"` // sdk 回傳的 token string
}
