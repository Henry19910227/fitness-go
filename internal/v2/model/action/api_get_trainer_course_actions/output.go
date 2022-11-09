package api_get_trainer_course_actions

import (
	"github.com/Henry19910227/fitness-go/internal/v2/field/action/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
)

// Output /v2/trainer/course/{course_id}/actions [GET] 獲取教練課表動作庫 API
type Output struct {
	base.Output
	Data   *Data          `json:"data,omitempty"`
	Paging *paging.Output `json:"paging,omitempty"`
}
type Data []*struct {
	optional.IDField
	optional.NameField
	optional.SourceField
	optional.TypeField
	optional.CategoryField
	optional.BodyField
	optional.EquipmentField
	optional.IntroField
	optional.CoverField
	optional.VideoField
	optional.StatusField
	optional.CreateAtField
	optional.UpdateAtField
}
