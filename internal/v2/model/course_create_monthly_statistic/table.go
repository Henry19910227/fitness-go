package course_create_monthly_statistic

import "github.com/Henry19910227/fitness-go/internal/v2/field/course_create_monthly_statistic/optional"

type Table struct {
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

func (Table) TableName() string {
	return "course_create_monthly_statistics"
}
