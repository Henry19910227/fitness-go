package workout_set

import (
	courseRequired "github.com/Henry19910227/fitness-go/internal/v2/field/course/required"
	"github.com/Henry19910227/fitness-go/internal/v2/field/workout_set/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/field/workout_set/required"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	"github.com/Henry19910227/fitness-go/internal/v2/model/file"
	orderBy "github.com/Henry19910227/fitness-go/internal/v2/model/order_by"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"github.com/Henry19910227/fitness-go/internal/v2/model/preload"
)

type PagingInput = paging.Input
type PreloadInput = preload.Input
type OrderByInput = orderBy.Input

// GenerateInput Test Input
type GenerateInput struct {
	DataAmount int
	WorkoutID  []*base.GenerateSetting
}

type FindInput struct {
	optional.IDField
	PreloadInput
}

type ListInput struct {
	optional.WorkoutIDField
	optional.TypeField
	PagingInput
	PreloadInput
}

type DeleteInput struct {
	required.IDField
}

type APIGetCMSWorkoutSetsInput struct {
	optional.WorkoutIDField
	PagingInput
}

// APICreateUserWorkoutSetsInput /v2/user/workout/{workout_id}/workout_sets [POST]
type APICreateUserWorkoutSetsInput struct {
	courseRequired.UserIDField
	Uri  APICreateUserWorkoutSetsUri
	Body APICreateUserWorkoutSetsBody
}
type APICreateUserWorkoutSetsUri struct {
	required.WorkoutIDField
}
type APICreateUserWorkoutSetsBody struct {
	ActionIDs []int64 `json:"action_ids" binding:"required,workout_set_action_ids" example:"1,10,15"` // 動作id
}

// APICreateUserWorkoutSetByDuplicateInput /v2/user/workout_set/{workout_set_id}/duplicate [POST]
type APICreateUserWorkoutSetByDuplicateInput struct {
	courseRequired.UserIDField
	Uri  APICreateUserWorkoutSetByDuplicateUri
	Body APICreateUserWorkoutSetByDuplicateBody
}
type APICreateUserWorkoutSetByDuplicateUri struct {
	required.IDField
}
type APICreateUserWorkoutSetByDuplicateBody struct {
	DuplicateCount int `json:"duplicate_count" binding:"required,min=1,max=5" example:"1"` //複製個數
}

// APICreateUserRestSetInput /v2/user/workout/{workout_id}/rest_set [POST]
type APICreateUserRestSetInput struct {
	courseRequired.UserIDField
	Uri APICreateUserRestSetUri
}
type APICreateUserRestSetUri struct {
	required.WorkoutIDField
}

// APIDeleteUserWorkoutSetInput /v2/user/workout_set/{workout_set_id} [POST]
type APIDeleteUserWorkoutSetInput struct {
	courseRequired.UserIDField
	Uri APIDeleteUserWorkoutSetUri
}
type APIDeleteUserWorkoutSetUri struct {
	required.IDField
}

// APIGetUserWorkoutSetsInput /v2/user/workout/{workout_id}/workout_sets [GET]
type APIGetUserWorkoutSetsInput struct {
	courseRequired.UserIDField
	Uri APIGetUserWorkoutSetsUri
}
type APIGetUserWorkoutSetsUri struct {
	required.WorkoutIDField
}

// APIUpdateUserWorkoutSetInput /v2/user/workout_set/{workout_set_id} [PATCH]
type APIUpdateUserWorkoutSetInput struct {
	courseRequired.UserIDField
	Uri  APIUpdateUserWorkoutSetUri
	Form APIUpdateUserWorkoutSetForm
}
type APIUpdateUserWorkoutSetUri struct {
	required.IDField
}
type APIUpdateUserWorkoutSetForm struct {
	optional.AutoNextField
	optional.RemarkField
	optional.WeightField
	optional.RepsField
	optional.DistanceField
	optional.DurationField
	optional.InclineField
	StartAudio    *file.Input
	ProgressAudio *file.Input
}

// APIDeleteUserWorkoutSetStartAudioInput /v2/user/workout_set/{workout_set_id}/start_audio [DELETE]
type APIDeleteUserWorkoutSetStartAudioInput struct {
	courseRequired.UserIDField
	Uri APIDeleteUserWorkoutSetStartAudioUri
}
type APIDeleteUserWorkoutSetStartAudioUri struct {
	required.IDField
}

// APIDeleteUserWorkoutSetProgressAudioInput /v2/user/workout_set/{workout_set_id}/progress_audio [DELETE]
type APIDeleteUserWorkoutSetProgressAudioInput struct {
	courseRequired.UserIDField
	Uri APIDeleteUserWorkoutSetProgressAudioUri
}
type APIDeleteUserWorkoutSetProgressAudioUri struct {
	required.IDField
}
