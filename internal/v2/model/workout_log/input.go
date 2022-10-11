package workout_log

import (
	"github.com/Henry19910227/fitness-go/internal/v2/field/workout_log/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/field/workout_log/required"
	"github.com/Henry19910227/fitness-go/internal/v2/model/workout_set_log"
)

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
