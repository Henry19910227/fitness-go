package workout_set_log

import "github.com/Henry19910227/fitness-go/internal/v2/model/workout_set_log/api_get_user_action_workout_set_logs"

type Resolver interface {
	APIGetUserActionWorkoutSetLogs(input *api_get_user_action_workout_set_logs.Input) (output api_get_user_action_workout_set_logs.Output)
}
