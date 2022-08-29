package workout_set

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
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
	IDOptional
}

type ListInput struct {
	WorkoutIDOptional
	TypeOptional
	PagingInput
	PreloadInput
}

type DeleteInput struct {
	IDRequired
}

type APIGetCMSWorkoutSetsInput struct {
	WorkoutIDField
	PagingInput
}

// APICreateUserWorkoutSetsInput /v2/user/workout/{workout_id}/workout_sets [POST]
type APICreateUserWorkoutSetsInput struct {
	UserID int64 `json:"user_id" binding:"required" example:"10001"` // 用戶 id
	Uri    APICreateUserWorkoutSetsUri
	Body   APICreateUserWorkoutSetsBody
}
type APICreateUserWorkoutSetsUri struct {
	WorkoutIDRequired
}
type APICreateUserWorkoutSetsBody struct {
	ActionIDs []int64 `json:"action_ids" binding:"required,workout_set_action_ids" example:"1,10,15"` // 動作id
}

// APIDeleteUserWorkoutSetInput /v2/user/workout_set_is/{workout_set_id} [POST]
type APIDeleteUserWorkoutSetInput struct {
	UserID int64 `json:"user_id" binding:"required" example:"10001"` // 用戶 id
	Uri    APIDeleteUserWorkoutSetUri
}
type APIDeleteUserWorkoutSetUri struct {
	IDRequired
}
