package user

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	"github.com/Henry19910227/fitness-go/internal/v2/model/trainer"
	"github.com/Henry19910227/fitness-go/internal/v2/model/user_subscribe_info"
)

type Output struct {
	Table
	Trainer           *trainer.Output             `json:"trainer,omitempty" gorm:"foreignKey:user_id;references:id"`             // 教練
	UserSubscribeInfo *user_subscribe_info.Output `json:"user_subscribe_info,omitempty" gorm:"foreignKey:user_id;references:id"` // 訂閱資訊
}

func (Output) TableName() string {
	return "users"
}

// APIUpdatePasswordOutput /v2/password [PATCH]
type APIUpdatePasswordOutput struct {
	base.Output
}

// APIUpdateUserProfileOutput /v2/user/profile [PATCH]
type APIUpdateUserProfileOutput struct {
	base.Output
}

// APIUpdateUserAvatarOutput /v2/user/avatar [PATCH]
type APIUpdateUserAvatarOutput struct {
	base.Output
	Data *string `json:"data,omitempty" example:"123.jpg"`
}

// APIGetUserProfileOutput /v2/user/profile [GET]
type APIGetUserProfileOutput struct {
	base.Output
	Data *APIGetUserProfileData `json:"data,omitempty"`
}
type APIGetUserProfileData struct {
	IDField
	AccountTypeField
	AccountField
	DeviceTokenField
	UserStatusField
	UserTypeField
	EmailField
	NicknameField
	AvatarField
	SexField
	BirthdayField
	HeightField
	WeightField
	ExperienceField
	TargetField
	CreateAtField
	UpdateAtField
}

// APIGetAppleRefreshTokenOutput /v2/apple_refresh_token [POST]
type APIGetAppleRefreshTokenOutput struct {
	base.Output
	Data *APIGetAppleRefreshTokenData `json:"data,omitempty"`
}
type APIGetAppleRefreshTokenData struct {
	RefreshToken string `json:"refresh_token" example:"EAAucgU8qZCzMBAOZCy59TLD1aM2NAO1ITBpZC64imFp95CRuPv4ZAWepAMV"` // 用於apple註冊與登入的刷新令牌
}

// APIRegisterForEmailOutput /v2/register/email [POST]
type APIRegisterForEmailOutput struct {
	base.Output
}

// APIRegisterForFacebookOutput /v2/register/facebook [POST]
type APIRegisterForFacebookOutput struct {
	base.Output
}

// APIRegisterForGoogleOutput /v2/register/google [POST]
type APIRegisterForGoogleOutput struct {
	base.Output
}

// APIRegisterForAppleOutput /v2/register/apple [POST]
type APIRegisterForAppleOutput struct {
	base.Output
}

// APIRegisterForLineOutput /v2/register/line [POST]
type APIRegisterForLineOutput struct {
	base.Output
}

// APILoginForEmailOutput /v2/login/email [POST]
type APILoginForEmailOutput struct {
	base.Output
	Data  *APILoginForEmailData `json:"data,omitempty"`
	Token *string               `json:"token,omitempty" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6I"` // Token

}
type APILoginForEmailData struct {
	IDField
	NicknameField
	AvatarField
	UserStatusField
	SexField
	BirthdayField
	HeightField
	WeightField
	ExperienceField
	TargetField
	CreateAtField
	UpdateAtField
	Trainer *struct {
		trainer.NicknameField
		trainer.AvatarField
		trainer.TrainerStatusField
		trainer.TrainerLevelField
		trainer.CreateAtField
		trainer.UpdateAtField
	} `json:"trainer,omitempty"`
	UserSubscribeInfo *struct {
		user_subscribe_info.StatusField
		user_subscribe_info.StartDateField
		user_subscribe_info.ExpiresDateField
		user_subscribe_info.UpdateAtField
	} `json:"user_subscribe_info,omitempty"`
}

// APILoginForFacebookOutput /v2/login/facebook [POST]
type APILoginForFacebookOutput struct {
	base.Output
	Data  *APILoginForFacebookData `json:"data,omitempty"`
	Token *string                  `json:"token,omitempty" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6I"` // Token

}
type APILoginForFacebookData struct {
	APILoginForEmailData
}

// APILoginForGoogleOutput /v2/login/google [POST]
type APILoginForGoogleOutput struct {
	base.Output
	Data  *APILoginForGoogleData `json:"data,omitempty"`
	Token *string                `json:"token,omitempty" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6I"` // Token

}
type APILoginForGoogleData struct {
	APILoginForEmailData
}

// APILoginForAppleOutput /v2/login/apple [POST]
type APILoginForAppleOutput struct {
	base.Output
	Data  *APILoginForAppleData `json:"data,omitempty"`
	Token *string                `json:"token,omitempty" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6I"` // Token

}
type APILoginForAppleData struct {
	APILoginForEmailData
}

// APILoginForLineOutput /v2/login/line [POST]
type APILoginForLineOutput struct {
	base.Output
	Data  *APILoginForLineData `json:"data,omitempty"`
	Token *string              `json:"token,omitempty" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6I"` // Token

}
type APILoginForLineData struct {
	APILoginForEmailData
}

// APILogoutOutput /v2/logout [POST]
type APILogoutOutput struct {
	base.Output
}

// APICreateRegisterOTPOutput /v2/otp [POST]
type APICreateRegisterOTPOutput struct {
	base.Output
	Data *APICreateRegisterOTPData `json:"data,omitempty"`
}
type APICreateRegisterOTPData struct {
	Code string `json:"otp_code" example:"254235"` // 信箱驗證碼
}

// APIRegisterNicknameValidateOutput /v2/register/nickname/validate [POST]
type APIRegisterNicknameValidateOutput struct {
	base.Output
}

// APIRegisterEmailValidateOutput /v2/register/email/validate [POST]
type APIRegisterEmailValidateOutput struct {
	base.Output
}

// APIRegisterEmailAccountValidateOutput /v2/register/email_account/validate [POST]
type APIRegisterEmailAccountValidateOutput struct {
	base.Output
}

// APIRegisterFacebookAccountValidateOutput /v2/register/facebook_account/validate [POST]
type APIRegisterFacebookAccountValidateOutput struct {
	base.Output
}

// APIRegisterLineAccountValidateOutput /v2/register/line_account/validate [POST]
type APIRegisterLineAccountValidateOutput struct {
	base.Output
}

// APIRegisterGoogleAccountValidateOutput /v2/google/google_account/validate [POST]
type APIRegisterGoogleAccountValidateOutput struct {
	base.Output
}

// APIRegisterAppleAccountValidateOutput /v2/google/apple_account/validate [POST]
type APIRegisterAppleAccountValidateOutput struct {
	base.Output
}
