package user_subscribe_info

import (
	"github.com/Henry19910227/fitness-go/internal/v2/field/user_subscribe_info/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
)

type Output struct {
	Table
}

func (Output) TableName() string {
	return "user_subscribe_infos"
}

// APIGetUserSubscribeInfoOutput /v2/user/subscribe_info [GET]
type APIGetUserSubscribeInfoOutput struct {
	base.Output
	Data *APIGetUserSubscribeInfoData `json:"data,omitempty"`
}
type APIGetUserSubscribeInfoData struct {
	optional.StatusField
	optional.StartDateField
	optional.ExpiresDateField
	optional.UpdateAtField
}
