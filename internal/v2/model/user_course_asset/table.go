package user_course_asset

import "github.com/Henry19910227/fitness-go/internal/v2/field/user_course_asset/optional"

type Table struct {
	optional.IDField
	optional.UserIDField
	optional.CourseIDField
	optional.AvailableField
	optional.CreateAtField
	optional.UpdateAtField
}
func (Table) TableName() string {
	return "user_course_assets"
}
