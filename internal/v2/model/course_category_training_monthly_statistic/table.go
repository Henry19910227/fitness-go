package course_category_training_monthly_statistic

import "github.com/Henry19910227/fitness-go/internal/v2/field/course_category_training_monthly_statistic/optional"

type Table struct {
	optional.CategoryField
	optional.YearField
	optional.MonthField
	optional.TotalField
	optional.MaleField
	optional.FemaleField
	optional.Age13to17Field
	optional.Age18to24Field
	optional.Age25to34Field
	optional.Age35to44Field
	optional.Age45to54Field
	optional.Age55to64Field
	optional.Age65UpField
	optional.CreateAtField
	optional.UpdateAtField
}

func (Table) TableName() string {
	return "course_category_training_monthly_statistics"
}
