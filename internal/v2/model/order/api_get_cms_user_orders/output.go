package api_get_cms_user_orders

import (
	courseOptional "github.com/Henry19910227/fitness-go/internal/v2/field/course/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/field/order/optional"
	trainerOptional "github.com/Henry19910227/fitness-go/internal/v2/field/trainer/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
)

// Output /v2/cms/user/{user_id}/orders [GET]
type Output struct {
	base.Output
	Data   *Data          `json:"data,omitempty"`
	Paging *paging.Output `json:"paging,omitempty"`
}
type Data []*struct {
	optional.IDField
	optional.CreateAtField
	OrderCourse *struct {
		Course *struct {
			courseOptional.IDField
			courseOptional.NameField
			Trainer *struct {
				trainerOptional.UserIDField
				trainerOptional.NicknameField
			} `json:"trainer,omitempty"`
		} `json:"course,omitempty"`
	} `json:"order_course,omitempty"`
}
