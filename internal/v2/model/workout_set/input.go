package workout_set

import (
	pagingOptional "github.com/Henry19910227/fitness-go/internal/v2/field/paging/optional"
	userRequired "github.com/Henry19910227/fitness-go/internal/v2/field/user/required"
	"github.com/Henry19910227/fitness-go/internal/v2/field/workout_set/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/field/workout_set/required"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	"github.com/Henry19910227/fitness-go/internal/v2/model/file"
	"github.com/Henry19910227/fitness-go/internal/v2/model/join"
	orderBy "github.com/Henry19910227/fitness-go/internal/v2/model/order_by"
	"github.com/Henry19910227/fitness-go/internal/v2/model/preload"
	"github.com/Henry19910227/fitness-go/internal/v2/model/where"
)

type PagingInput = struct {
	pagingOptional.PageField
	pagingOptional.SizeField
}
type PreloadInput = preload.Input
type WhereInput = where.Input
type JoinInput = join.Input
type OrderByInput = orderBy.Input
type CustomOrderByInput = orderBy.CustomInput

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
	optional.ActionIDField
	optional.TypeField
	JoinInput
	WhereInput
	PagingInput
	PreloadInput
	OrderByInput
	CustomOrderByInput
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
	userRequired.UserIDField
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
	userRequired.UserIDField
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
	userRequired.UserIDField
	Uri APICreateUserRestSetUri
}
type APICreateUserRestSetUri struct {
	required.WorkoutIDField
}

// APIDeleteUserWorkoutSetInput /v2/user/workout_set/{workout_set_id} [POST]
type APIDeleteUserWorkoutSetInput struct {
	userRequired.UserIDField
	Uri APIDeleteUserWorkoutSetUri
}
type APIDeleteUserWorkoutSetUri struct {
	required.IDField
}

// APIGetUserWorkoutSetsInput /v2/user/workout/{workout_id}/workout_sets [GET]
type APIGetUserWorkoutSetsInput struct {
	userRequired.UserIDField
	Uri APIGetUserWorkoutSetsUri
}
type APIGetUserWorkoutSetsUri struct {
	required.WorkoutIDField
}

// APIUpdateUserWorkoutSetInput /v2/user/workout_set/{workout_set_id} [PATCH]
type APIUpdateUserWorkoutSetInput struct {
	userRequired.UserIDField
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
	userRequired.UserIDField
	Uri APIDeleteUserWorkoutSetStartAudioUri
}
type APIDeleteUserWorkoutSetStartAudioUri struct {
	required.IDField
}

// APIDeleteUserWorkoutSetProgressAudioInput /v2/user/workout_set/{workout_set_id}/progress_audio [DELETE]
type APIDeleteUserWorkoutSetProgressAudioInput struct {
	userRequired.UserIDField
	Uri APIDeleteUserWorkoutSetProgressAudioUri
}
type APIDeleteUserWorkoutSetProgressAudioUri struct {
	required.IDField
}

// APIGetTrainerWorkoutSetsInput /v2/trainer/workout/{workout_id}/workout_sets [GET]
type APIGetTrainerWorkoutSetsInput struct {
	userRequired.UserIDField
	Uri APIGetUserWorkoutSetsUri
}
type APIGetTrainerWorkoutSetsUri struct {
	required.WorkoutIDField
}

// APICreateTrainerWorkoutSetsInput /v2/trainer/workout/{workout_id}/workout_sets [POST]
type APICreateTrainerWorkoutSetsInput struct {
	userRequired.UserIDField
	Uri  APICreateTrainerWorkoutSetsUri
	Body APICreateTrainerWorkoutSetsBody
}
type APICreateTrainerWorkoutSetsUri struct {
	required.WorkoutIDField
}
type APICreateTrainerWorkoutSetsBody struct {
	ActionIDs []int64 `json:"action_ids" binding:"required,workout_set_action_ids" example:"1,10,15"` // 動作id
}

// APICreateTrainerRestSetInput /v2/trainer/workout/{workout_id}/rest_set [POST]
type APICreateTrainerRestSetInput struct {
	userRequired.UserIDField
	Uri APICreateTrainerRestSetUri
}
type APICreateTrainerRestSetUri struct {
	required.WorkoutIDField
}

// APIDeleteTrainerWorkoutSetInput /v2/trainer/workout_set/{workout_set_id} [DELETE]
type APIDeleteTrainerWorkoutSetInput struct {
	userRequired.UserIDField
	Uri APIDeleteTrainerWorkoutSetUri
}
type APIDeleteTrainerWorkoutSetUri struct {
	required.IDField
}

// APIUpdateTrainerWorkoutSetInput /v2/trainer/workout_set/{workout_set_id} [PATCH]
type APIUpdateTrainerWorkoutSetInput struct {
	userRequired.UserIDField
	Uri  APIUpdateTrainerWorkoutSetUri
	Form APIUpdateTrainerWorkoutSetForm
}
type APIUpdateTrainerWorkoutSetUri struct {
	required.IDField
}
type APIUpdateTrainerWorkoutSetForm struct {
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

// APIDeleteTrainerWorkoutSetStartAudioInput /v2/trainer/workout_set/{workout_set_id}/start_audio [DELETE]
type APIDeleteTrainerWorkoutSetStartAudioInput struct {
	userRequired.UserIDField
	Uri APIDeleteTrainerWorkoutSetStartAudioUri
}
type APIDeleteTrainerWorkoutSetStartAudioUri struct {
	required.IDField
}

// APIDeleteTrainerWorkoutSetProgressAudioInput /v2/trainer/workout_set/{workout_set_id}/progress_audio [DELETE]
type APIDeleteTrainerWorkoutSetProgressAudioInput struct {
	userRequired.UserIDField
	Uri APIDeleteTrainerWorkoutSetProgressAudioUri
}
type APIDeleteTrainerWorkoutSetProgressAudioUri struct {
	required.IDField
}

// APICreateTrainerWorkoutSetByDuplicateInput /v2/trainer/workout_set/{workout_set_id}/duplicate [POST]
type APICreateTrainerWorkoutSetByDuplicateInput struct {
	userRequired.UserIDField
	Uri  APICreateTrainerWorkoutSetByDuplicateUri
	Body APICreateTrainerWorkoutSetByDuplicateBody
}
type APICreateTrainerWorkoutSetByDuplicateUri struct {
	required.IDField
}
type APICreateTrainerWorkoutSetByDuplicateBody struct {
	DuplicateCount int `json:"duplicate_count" binding:"required,min=1,max=5" example:"1"` //複製個數
}

// APIGetStoreWorkoutSetsInput /v2/store/workout/{workout_id}/workout_sets [GET]
type APIGetStoreWorkoutSetsInput struct {
	userRequired.UserIDField
	Uri APIGetStoreWorkoutSetsUri
}
type APIGetStoreWorkoutSetsUri struct {
	required.WorkoutIDField
}
