package workout

import (
	userRequired "github.com/Henry19910227/fitness-go/internal/v2/field/user/required"
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
	userRequired.UserIDField
	Uri  APICreateUserWorkoutUri
	Body APICreateUserWorkoutBody
}
type APICreateUserWorkoutUri struct {
	required.PlanIDField
}
type APICreateUserWorkoutBody struct {
	required.NameField
	WorkoutTemplateID *int64 `json:"workout_template_id" binding:"omitempty" example:"1"` // 訓練模板ID
}

// APIDeleteUserWorkoutInput /v2/user/workout/{workout_id} [DELETE]
type APIDeleteUserWorkoutInput struct {
	userRequired.UserIDField
	Uri APIDeleteUserWorkoutUri
}
type APIDeleteUserWorkoutUri struct {
	required.IDField
}

// APIGetUserWorkoutsInput /v2/user/plan/{plan_is}/workouts [GET]
type APIGetUserWorkoutsInput struct {
	userRequired.UserIDField
	Uri APIGetUserPlansUri
}
type APIGetUserPlansUri struct {
	required.PlanIDField
}

// APIUpdateUserWorkoutInput /v2/user/workout/{workout_id} [PATCH]
type APIUpdateUserWorkoutInput struct {
	userRequired.UserIDField
	Uri  APIDeleteUserWorkoutUri
	Form APIUpdateUserWorkoutForm
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
	userRequired.UserIDField
	Uri APIDeleteUserWorkoutStartAudioUri
}
type APIDeleteUserWorkoutStartAudioUri struct {
	required.IDField
}

// APIDeleteUserWorkoutEndAudioInput /v2/user/workout/{workout_id}/end_audio [DELETE]
type APIDeleteUserWorkoutEndAudioInput struct {
	userRequired.UserIDField
	Uri APIDeleteUserWorkoutEndAudioUri
}
type APIDeleteUserWorkoutEndAudioUri struct {
	required.IDField
}

// APIGetTrainerWorkoutsInput /v2/trainer/plan/{plan_is}/workouts [GET]
type APIGetTrainerWorkoutsInput struct {
	userRequired.UserIDField
	Uri APIGetTrainerPlansUri
}
type APIGetTrainerPlansUri struct {
	required.PlanIDField
}

// APICreateTrainerWorkoutInput /v2/trainer/plan/{plan_id}/workout [POST]
type APICreateTrainerWorkoutInput struct {
	userRequired.UserIDField
	Uri  APICreateTrainerWorkoutUri
	Body APICreateTrainerWorkoutBody
}
type APICreateTrainerWorkoutUri struct {
	required.PlanIDField
}
type APICreateTrainerWorkoutBody struct {
	required.NameField
	WorkoutTemplateID *int64 `json:"workout_template_id" binding:"omitempty" example:"1"` // 訓練模板ID
}

// APIUpdateTrainerWorkoutInput /v2/trainer/workout/{workout_id} [PATCH]
type APIUpdateTrainerWorkoutInput struct {
	userRequired.UserIDField
	Uri  APIUpdateTrainerWorkoutUri
	Form APIUpdateTrainerWorkoutForm
}
type APIUpdateTrainerWorkoutUri struct {
	required.IDField
}
type APIUpdateTrainerWorkoutForm struct {
	optional.EquipmentField
	optional.NameField
	StartAudio *file.Input
	EndAudio   *file.Input
}

// APIDeleteTrainerWorkoutInput /v2/trainer/workout/{workout_id} [DELETE]
type APIDeleteTrainerWorkoutInput struct {
	userRequired.UserIDField
	Uri APIDeleteTrainerWorkoutUri
}
type APIDeleteTrainerWorkoutUri struct {
	required.IDField
}

// APIDeleteTrainerWorkoutStartAudioInput /v2/trainer/workout/{workout_id}/start_audio [DELETE]
type APIDeleteTrainerWorkoutStartAudioInput struct {
	userRequired.UserIDField
	Uri APIDeleteTrainerWorkoutStartAudioUri
}
type APIDeleteTrainerWorkoutStartAudioUri struct {
	required.IDField
}
