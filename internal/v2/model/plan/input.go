package plan

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	orderBy "github.com/Henry19910227/fitness-go/internal/v2/model/order_by"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"github.com/Henry19910227/fitness-go/internal/v2/model/preload"
)

type PagingInput = paging.Input
type PreloadInput = preload.Input
type OrderByInput = orderBy.Input

type GenerateInput struct {
	DataAmount int
	CourseID   []*base.GenerateSetting
}

type FindInput struct {
	IDField
	WorkoutID *int64 `json:"workout_id,omitempty"` // 訓練 id
}

type DeleteInput struct {
	IDRequired
}

type ListInput struct {
	CourseIDField
	PagingInput
	OrderByInput
	PreloadInput
}

type APIGetCMSPlansInput struct {
	CourseIDField
	PagingInput
	OrderByInput
}

// APICreatePersonalPlanInput /v2/personal/course/{course_id}/plan [POST]
type APICreatePersonalPlanInput struct {
	UserID int64 `json:"user_id" binding:"required" example:"10001"` // 用戶 id
	Uri    APICreatePersonalPlanUri
	Body   APICreatePersonalPlanBody
}
type APICreatePersonalPlanUri struct {
	CourseIDRequired
}
type APICreatePersonalPlanBody struct {
	NameRequired
}

// APIDeletePersonalPlanInput /v2/personal/workout/{workout_id} [DELETE]
type APIDeletePersonalPlanInput struct {
	UserID int64 `json:"user_id" binding:"required" example:"10001"` // 用戶 id
	Uri    APIDeletePersonalPlanUri
}
type APIDeletePersonalPlanUri struct {
	IDRequired
}
