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

// APIRegisterAccountValidateInput /v2/register/account/validate [POST]
type APIRegisterAccountValidateInput struct {
	Body APIRegisterAccountValidateBody
}
type APIRegisterAccountValidateBody struct {
	EmailRequired
}
