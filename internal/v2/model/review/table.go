package review

import "github.com/Henry19910227/fitness-go/internal/v2/field/review/optional"

type Table struct {
	optional.IDField
	optional.CourseIDField
	optional.UserIDField
	optional.ScoreField
	optional.BodyField
	optional.CreateAtField
}

func (Table) TableName() string {
	return "reviews"
}
