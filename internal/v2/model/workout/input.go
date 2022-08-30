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

type FindInput struct {
	IDOptional
}

type DeleteInput struct {
	IDRequired
}

// APICreateUserWorkoutInput /v2/user/plan/{plan_id}/workout [POST]
type APICreateUserWorkoutInput struct {
	UserID int64 `json:"user_id" binding:"required" example:"10001"` // 用戶 id
	Uri    APICreateUserWorkoutUri
	Body   APICreateUserWorkoutBody
}
type APICreateUserWorkoutUri struct {
	PlanIDRequired
}
type APICreateUserWorkoutBody struct {
	NameRequired
}

// APIDeleteUserWorkoutInput /v2/user/workout/{workout_id} [DELETE]
type APIDeleteUserWorkoutInput struct {
	UserID int64 `json:"user_id" binding:"required" example:"10001"` // 用戶 id
	Uri    APIDeleteUserWorkoutUri
}
type APIDeleteUserWorkoutUri struct {
	IDRequired
}

// APIGetUserWorkoutsInput /v2/user/plan/{plan_is}/workouts [GET]
type APIGetUserWorkoutsInput struct {
	UserID int64 `json:"user_id" binding:"required" example:"10001"` // 用戶 id
	Uri    APIGetUserPlansUri
}
type APIGetUserPlansUri struct {
	PlanIDRequired
}
