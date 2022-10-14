package review_statistic

import "github.com/Henry19910227/fitness-go/internal/v2/field/review_statistic/optional"

type Table struct {
	optional.CourseIDField
	optional.ScoreTotalField
	optional.AmountField
	optional.FiveTotalField
	optional.FourTotalField
	optional.ThreeTotalField
	optional.TwoTotalField
	optional.OneTotalField
	optional.UpdateAtField
}

func (Table) TableName() string {
	return "review_statistics"
}
