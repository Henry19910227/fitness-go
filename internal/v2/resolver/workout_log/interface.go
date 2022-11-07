package workout_log

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/workout_log"
	"github.com/Henry19910227/fitness-go/internal/v2/model/workout_log/api_delete_user_workout_log"
	"github.com/Henry19910227/fitness-go/internal/v2/model/workout_log/api_get_user_workout_log"
	"gorm.io/gorm"
)

type Resolver interface {
	APICreateUserWorkoutLog(tx *gorm.DB, input *model.APICreateUserWorkoutLogInput) (output model.APICreateUserWorkoutLogOutput)
	APIGetUserWorkoutLogs(input *model.APIGetUserWorkoutLogsInput) (output model.APIGetUserWorkoutLogsOutput)
	APIGetUserWorkoutLog(input *api_get_user_workout_log.Input) (output api_get_user_workout_log.Output)
	APIDeleteUserWorkoutLog(input *api_delete_user_workout_log.Input) (output api_delete_user_workout_log.Output)
}
