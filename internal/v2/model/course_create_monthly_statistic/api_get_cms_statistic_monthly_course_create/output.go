package api_get_cms_statistic_monthly_course_create

import (
	"github.com/Henry19910227/fitness-go/internal/v2/field/course_create_monthly_statistic/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
)

// Output /v2/cms/statistic_monthly/course/create [GET]
type Output struct {
	base.Output
	Data *Data `json:"data,omitempty"`
}
type Data struct {
	optional.YearField
	optional.MonthField
	optional.TotalField
	optional.FreeField
	optional.SubscribeField
	optional.ChargeField
	optional.AerobicField
	optional.IntervalTrainingField
	optional.WeightTrainingField
	optional.ResistanceTrainingField
	optional.BodyweightTrainingField
	optional.OtherTrainingField
	optional.CreateAtField
	optional.UpdateAtField
}
