package plan

import (
	courseRequired "github.com/Henry19910227/fitness-go/internal/v2/field/course/required"
	"github.com/Henry19910227/fitness-go/internal/v2/field/plan/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/field/plan/required"
	userRequired "github.com/Henry19910227/fitness-go/internal/v2/field/user/required"
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
	optional.IDField
	WorkoutID *int64 `json:"workout_id,omitempty"` // 訓練 id
}

type DeleteInput struct {
	required.IDField
}

type ListInput struct {
	optional.CourseIDField
	PagingInput
	OrderByInput
	PreloadInput
}

type APIGetCMSPlansInput struct {
	Uri   APIGetCMSPlansUri
	Query APIGetCMSPlansQuery
}
type APIGetCMSPlansUri struct {
	required.CourseIDField
}
type APIGetCMSPlansQuery struct {
	PagingInput
	OrderByInput
}

// APICreateUserPlanInput /v2/user/course/{course_id}/plan [POST]
type APICreateUserPlanInput struct {
	courseRequired.UserIDField
	Uri  APICreateUserPlanUri
	Body APICreateUserPlanBody
}
type APICreateUserPlanUri struct {
	required.CourseIDField
}
type APICreateUserPlanBody struct {
	required.NameField
}

// APIDeleteUserPlanInput /v2/user/workout/{workout_id} [DELETE]
type APIDeleteUserPlanInput struct {
	userRequired.UserIDField
	Uri APIDeleteUserPlanUri
}
type APIDeleteUserPlanUri struct {
	required.IDField
}

// APIGetUserPlansInput /v2/user/course/{course_id}/plans [GET]
type APIGetUserPlansInput struct {
	userRequired.UserIDField
	Uri APIGetUserPlansUri
}
type APIGetUserPlansUri struct {
	required.CourseIDField
}

// APIUpdateUserPlanInput /v2/user/plan/{plan_id} [PATCH]
type APIUpdateUserPlanInput struct {
	userRequired.UserIDField
	Uri  APIUpdateUserPlanUri
	Body APIUpdateUserPlanBody
}
type APIUpdateUserPlanUri struct {
	required.IDField
}
type APIUpdateUserPlanBody struct {
	required.NameField
}

// APICreateTrainerPlanInput /v2/trainer/course/{course_id}/plan [POST]
type APICreateTrainerPlanInput struct {
	userRequired.UserIDField
	Uri  APICreateTrainerPlanUri
	Body APICreateTrainerPlanBody
}
type APICreateTrainerPlanUri struct {
	required.CourseIDField
}
type APICreateTrainerPlanBody struct {
	required.NameField
}

// APIGetTrainerPlansInput /v2/trainer/course/{course_id}/plans [GET]
type APIGetTrainerPlansInput struct {
	userRequired.UserIDField
	Uri APIGetTrainerPlansUri
}
type APIGetTrainerPlansUri struct {
	required.CourseIDField
}

// APIGetProductPlansInput /v2/product/course/{course_id}/plans [GET]
type APIGetProductPlansInput struct {
	userRequired.UserIDField
	Uri APIGetProductPlansUri
}
type APIGetProductPlansUri struct {
	required.CourseIDField
}
