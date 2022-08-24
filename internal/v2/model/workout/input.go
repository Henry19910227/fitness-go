package workout

import (
	orderBy "github.com/Henry19910227/fitness-go/internal/v2/model/order_by"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"github.com/Henry19910227/fitness-go/internal/v2/model/preload"
)

type PagingInput = paging.Input
type PreloadInput = preload.Input
type OrderByInput = orderBy.Input

type ListInput struct {
	CourseID *int64 `json:"course_id,omitempty" gorm:"column:course_id" example:"10"` //課表id
	PlanIDField
	PagingInput
	OrderByInput
	PreloadInput
}

// APICreatePersonalWorkoutInput /v2/personal/plan/{plan_id}/workout [POST]
type APICreatePersonalWorkoutInput struct {
	UserID int64 `json:"user_id" binding:"required" example:"10001"` // 用戶 id
	Uri    APICreatePersonalWorkoutUri
	Body   APICreatePersonalWorkoutBody
}
type APICreatePersonalWorkoutUri struct {
	PlanIDRequired
}
type APICreatePersonalWorkoutBody struct {
	NameRequired
}
