package api_get_cms_trainer_courses

import (
	courseRequired "github.com/Henry19910227/fitness-go/internal/v2/field/course/required"
	orderByOptional "github.com/Henry19910227/fitness-go/internal/v2/field/order_by/optional"
	pagingOptional "github.com/Henry19910227/fitness-go/internal/v2/field/paging/optional"
)

// Input /v2/cms/trainer/{user_id}/courses [GET]
type Input struct {
	Uri   Uri
	Query Query
}
type Uri struct {
	courseRequired.UserIDField
}
type Query struct {
	orderByOptional.OrderFieldField
	orderByOptional.OrderTypeField
	pagingOptional.PageField
	pagingOptional.SizeField
}
