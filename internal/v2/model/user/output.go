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

// APIRegisterEmailOutput /v2/register/email [POST]
type APIRegisterEmailOutput struct {
	base.Output
	Data *APIRegisterEmailData `json:"data,omitempty"`
}
type APIRegisterEmailData struct {
	IDField
}
