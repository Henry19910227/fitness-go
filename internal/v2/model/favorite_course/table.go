package favorite_course

import "github.com/Henry19910227/fitness-go/internal/v2/field/favorite_course/optional"

type Table struct {
	optional.UserIDField
	optional.CourseIDField
	optional.CreateAtField
}
func (Table) TableName() string {
	return "favorite_courses"
}
