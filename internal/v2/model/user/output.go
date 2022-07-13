package user

import "github.com/Henry19910227/fitness-go/internal/v2/model/base"

type Output struct {
	Table
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
	Data *APIRegisterForEmailData `json:"data,omitempty"`
}
type APIRegisterForEmailData struct {
	IDField
}

// APIRegisterNicknameValidateOutput /v2/register/nickname/validate [POST]
type APIRegisterNicknameValidateOutput struct {
	base.Output
}

// APIRegisterAccountValidateOutput /v2/register/account/validate [POST]
type APIRegisterAccountValidateOutput struct {
	base.Output
}
