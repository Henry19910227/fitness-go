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

// APIRegisterForEmailOutput /v2/register/email [POST]
type APIRegisterForEmailOutput struct {
	base.Output
}

// APILoginForEmailOutput /v2/login/email [POST]
type APILoginForEmailOutput struct {
	base.Output
	Data  *APILoginForEmailData `json:"data,omitempty"`
	Token string                `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6I"` // Token

}
type APILoginForEmailData struct {
	IDField
	NicknameField
	AvatarField
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
		user_subscribe_info.EndDateField
		user_subscribe_info.UpdateAtField
	} `json:"user_subscribe_info,omitempty"`
}

// APICreateRegisterOTPOutput /v2/register/otp [POST]
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

// APIRegisterAccountValidateOutput /v2/register/account/validate [POST]
type APIRegisterAccountValidateOutput struct {
	base.Output
}
