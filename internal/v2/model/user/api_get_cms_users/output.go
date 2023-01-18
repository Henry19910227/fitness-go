package api_get_cms_users

import (
	userOptional "github.com/Henry19910227/fitness-go/internal/v2/field/user/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
)

// Output /v2/cms/users [GET]
type Output struct {
	base.Output
	Data   *Data          `json:"data,omitempty"`
	Paging *paging.Output `json:"paging,omitempty"`
}
type Data []*struct {
	userOptional.IDField
	userOptional.AccountTypeField
	userOptional.UserStatusField
	userOptional.UserTypeField
	userOptional.EmailField
	userOptional.NicknameField
	userOptional.CreateAtField
	userOptional.UpdateAtField
}
