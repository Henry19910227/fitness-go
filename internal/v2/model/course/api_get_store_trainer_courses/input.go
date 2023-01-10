package api_get_store_trainer_courses

import (
	courseOptional "github.com/Henry19910227/fitness-go/internal/v2/field/course/optional"
	pagingOptional "github.com/Henry19910227/fitness-go/internal/v2/field/paging/optional"
	userRequired "github.com/Henry19910227/fitness-go/internal/v2/field/user/required"
)

// Input /v2/store/trainer/{user_id}/courses [GET]
type Input struct {
	userRequired.UserIDField
	Uri   Uri
	Query Query
}
type Uri struct {
	userRequired.UserIDField
}
type Query struct {
	courseOptional.SaleTypeField
	pagingOptional.PageField
	pagingOptional.SizeField
}
