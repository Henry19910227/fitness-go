package course_status_update_log

import "github.com/Henry19910227/fitness-go/internal/v2/field/course_status_update_log/optional"

type Table struct {
	optional.IDField
	optional.CourseIDField
	optional.CourseStatusField
	optional.CommentField
	optional.CreateAtField
}

func (Table) TableName() string {
	return "course_status_update_logs"
}
