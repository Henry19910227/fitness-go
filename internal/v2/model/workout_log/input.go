package workout_log

import (
	"github.com/Henry19910227/fitness-go/internal/v2/field/workout_log/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/field/workout_log/required"
	"github.com/Henry19910227/fitness-go/internal/v2/model/join"
	orderBy "github.com/Henry19910227/fitness-go/internal/v2/model/order_by"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"github.com/Henry19910227/fitness-go/internal/v2/model/preload"
	"github.com/Henry19910227/fitness-go/internal/v2/model/where"
	"github.com/Henry19910227/fitness-go/internal/v2/model/workout_set_log"
)

type PagingInput = paging.Input
type PreloadInput = preload.Input
type WhereInput = where.Input
type JoinInput = join.Input
type OrderByInput = orderBy.Input
type CustomOrderByInput = orderBy.CustomInput

type ListInput struct {
	optional.UserIDField
	JoinInput
	WhereInput
	PagingInput
	PreloadInput
	OrderByInput
	CustomOrderByInput
}

// APICreateUserWorkoutLogInput /v2/user/workout/{workout_id}/workout_log [POST]
type APICreateUserWorkoutLogInput struct {
	required.UserIDField
	Uri  APICreateUserWorkoutLogUri
	Body APICreateUserWorkoutLogBody
}
type APICreateUserWorkoutLogUri struct {
	required.WorkoutIDField
}
type APICreateUserWorkoutLogBody struct {
	required.DurationField
	optional.IntensityField
	optional.PlaceField
	WorkoutSetLogs []*struct {
		workout_set_log.WorkoutSetIDRequired
		workout_set_log.WeightRequired
		workout_set_log.DistanceRequired
		workout_set_log.InclineRequired
		workout_set_log.RepsRequired
		workout_set_log.DurationRequired
	} `json:"workout_set_logs"`
}

// APIGetUserWorkoutLogsInput /v2/user/workout_logs [GET]
type APIGetUserWorkoutLogsInput struct {
	required.UserIDField
	Query APIGetUserWorkoutLogsQuery
}
type APIGetUserWorkoutLogsQuery struct {
	StartDate string `form:"start_date" binding:"required,datetime=2006-01-02" example:"區間開始日期"` //區間開始日期
	EndDate   string `form:"end_date" binding:"required,datetime=2006-01-02" example:"區間結束日期"`   //區間結束日期
	PagingInput
}
