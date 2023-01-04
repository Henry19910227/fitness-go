package api_cms_login

import (
	"github.com/Henry19910227/fitness-go/internal/v2/field/admin/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
)

// Output /v2/cms/login [POST]
type Output struct {
	base.Output
	Data  *Data   `json:"data,omitempty"`
	Token *string `json:"token,omitempty" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6I"` // Token
}
type Data struct {
	optional.IDField
	optional.EmailField
	optional.NicknameField
	optional.LvField
}
