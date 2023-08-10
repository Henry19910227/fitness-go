package api_get_cms_trainers

import (
	orderByOptional "github.com/Henry19910227/fitness-go/internal/v2/field/order_by/optional"
	pagingOptional "github.com/Henry19910227/fitness-go/internal/v2/field/paging/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/field/trainer/optional"
)

// Input /v2/cms/trainer/{user_id} [PATCH]
type Input struct {
	Query Query
}
type Query struct {
	optional.UserIDField
	optional.NicknameField
	Email *string `json:"email,omitempty" form:"email" gorm:"column:email" binding:"omitempty,max=255" example:"henry@gmail.com"` // 信箱
	optional.TrainerStatusField
	orderByOptional.OrderFieldField
	orderByOptional.OrderTypeField
	pagingOptional.PageField
	pagingOptional.SizeField
}
