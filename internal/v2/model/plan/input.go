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

// APICreateUserPlanInput /v2/user/course/{course_id}/plan [POST]
type APICreateUserPlanInput struct {
	UserID int64 `json:"user_id" binding:"required" example:"10001"` // 用戶 id
	Uri    APICreateUserPlanUri
	Body   APICreateUserPlanBody
}
type APICreateUserPlanUri struct {
	CourseIDRequired
}
type APICreateUserPlanBody struct {
	NameRequired
}

// APIDeleteUserPlanInput /v2/user/workout/{workout_id} [DELETE]
type APIDeleteUserPlanInput struct {
	UserID int64 `json:"user_id" binding:"required" example:"10001"` // 用戶 id
	Uri    APIDeleteUserPlanUri
}
type APIDeleteUserPlanUri struct {
	IDRequired
}

// APIGetUserPlansInput /v2/user/course/{course_id}/plans [GET]
type APIGetUserPlansInput struct {
	UserID int64 `json:"user_id" binding:"required" example:"10001"` // 用戶 id
	Uri    APIGetUserPlansUri
}
type APIGetUserPlansUri struct {
	CourseIDRequired
}

// APIUpdateUserPlanInput /v2/user/plan/{plan_id} [PATCH]
type APIUpdateUserPlanInput struct {
	UserID int64 `json:"user_id" binding:"required" example:"10001"` // 用戶 id
	Uri    APIUpdateUserPlanUri
	Body   APIUpdateUserPlanBody
}
type APIUpdateUserPlanUri struct {
	IDRequired
}
type APIUpdateUserPlanBody struct {
	NameRequired
}

// APICreateTrainerPlanInput /v2/trainer/course/{course_id}/plan [POST]
type APICreateTrainerPlanInput struct {
	UserID int64 `json:"user_id" binding:"required" example:"10001"` // 用戶 id
	Uri    APICreateTrainerPlanUri
	Body   APICreateTrainerPlanBody
}
type APICreateTrainerPlanUri struct {
	CourseIDRequired
}
type APICreateTrainerPlanBody struct {
	NameRequired
}
