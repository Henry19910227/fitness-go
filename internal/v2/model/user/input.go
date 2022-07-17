package user

import (
	orderBy "github.com/Henry19910227/fitness-go/internal/v2/model/order_by"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"github.com/Henry19910227/fitness-go/internal/v2/model/preload"
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

// APILoginForEmailInput /v2/login/email [POST]
type APILoginForEmailInput struct {
	Body APILoginForEmailBody
}
type APILoginForEmailBody struct {
	EmailRequired
	PasswordRequired
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
	AccessToken string `json:"access_token" binding:"required" example:"EAAucgU8qZCzMBAOZCy59TLD1aM2NAO1ITBpZC64imFp95CRuPv4ZAWepAMV"` // facebook sdk 回傳的 token string
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
	AccessToken string `json:"access_token" binding:"required" example:"EAAucgU8qZCzMBAOZCy59TLD1aM2NAO1ITBpZC64imFp95CRuPv4ZAWepAMV"` // facebook sdk 回傳的 token string
}
