package api_get_cms_course_users

import (
	courseRequired "github.com/Henry19910227/fitness-go/internal/v2/field/course/required"
	orderByRequired "github.com/Henry19910227/fitness-go/internal/v2/field/order_by/required"
	pagingOptional "github.com/Henry19910227/fitness-go/internal/v2/field/paging/optional"
)

// Input /v2/cms/course/{course_id}/users [GET]
type Input struct {
	Uri   Uri
	Query Query
}
type Uri struct {
	courseRequired.CourseIDField
}
type Query struct {
	pagingOptional.PageField
	pagingOptional.SizeField
	orderByRequired.OrderFieldField
	orderByRequired.OrderTypeField
}
