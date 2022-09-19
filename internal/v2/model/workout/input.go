package workout

import (
	"github.com/Henry19910227/fitness-go/internal/v2/field/workout/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/field/workout/required"
	"github.com/Henry19910227/fitness-go/internal/v2/model/file"
	orderBy "github.com/Henry19910227/fitness-go/internal/v2/model/order_by"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"github.com/Henry19910227/fitness-go/internal/v2/model/preload"
)

type PagingInput = paging.Input
type PreloadInput = preload.Input
type OrderByInput = orderBy.Input

type ListInput struct {
	CourseID *int64 `json:"course_id,omitempty" gorm:"column:course_id" example:"10"` //課表id
	optional.PlanIDField
	PagingInput
	OrderByInput
	PreloadInput
}

type FindInput struct {
	optional.IDField
}

type DeleteInput struct {
	required.IDField
}

// APICreateUserWorkoutInput /v2/user/plan/{plan_id}/workout [POST]
type APICreateUserWorkoutInput struct {
	UserID int64 `json:"user_id" binding:"required" example:"10001"` // 用戶 id
	Uri    APICreateUserWorkoutUri
	Body   APICreateUserWorkoutBody
}
type APICreateUserWorkoutUri struct {
	required.PlanIDField
}
type APICreateUserWorkoutBody struct {
	required.NameField
}

// APIDeleteUserWorkoutInput /v2/user/workout/{workout_id} [DELETE]
type APIDeleteUserWorkoutInput struct {
	UserID int64 `json:"user_id" binding:"required" example:"10001"` // 用戶 id
	Uri    APIDeleteUserWorkoutUri
}
type APIDeleteUserWorkoutUri struct {
	required.IDField
}

// APIGetUserWorkoutsInput /v2/user/plan/{plan_is}/workouts [GET]
type APIGetUserWorkoutsInput struct {
	UserID int64 `json:"user_id" binding:"required" example:"10001"` // 用戶 id
	Uri    APIGetUserPlansUri
}
type APIGetUserPlansUri struct {
	required.PlanIDField
}

// APIUpdateUserWorkoutInput /v2/user/workout/{workout_id} [PATCH]
type APIUpdateUserWorkoutInput struct {
	UserID int64 `json:"user_id" binding:"required" example:"10001"` // 用戶 id
	Uri    APIDeleteUserWorkoutUri
	Form   APIUpdateUserWorkoutForm
}
type APIUpdateUserWorkoutUri struct {
	required.IDField
}
type APIUpdateUserWorkoutForm struct {
	optional.EquipmentField
	optional.NameField
	StartAudio *file.Input
	EndAudio   *file.Input
}

// APIDeleteUserWorkoutStartAudioInput /v2/user/workout/{workout_id}/start_audio [DELETE]
type APIDeleteUserWorkoutStartAudioInput struct {
	UserID int64 `json:"user_id" binding:"required" example:"10001"` // 用戶 id
	Uri    APIDeleteUserWorkoutStartAudioUri
}
type APIDeleteUserWorkoutStartAudioUri struct {
	required.IDField
}

// APIDeleteUserWorkoutEndAudioInput /v2/user/workout/{workout_id}/end_audio [DELETE]
type APIDeleteUserWorkoutEndAudioInput struct {
	UserID int64 `json:"user_id" binding:"required" example:"10001"` // 用戶 id
	Uri    APIDeleteUserWorkoutEndAudioUri
}
type APIDeleteUserWorkoutEndAudioUri struct {
	required.IDField
}
