package workout_log

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/workout_log"
	"gorm.io/gorm"
)

type Resolver interface {
	APICreateUserWorkoutLog(tx *gorm.DB, input *model.APICreateUserWorkoutLogInput) (output model.APICreateUserWorkoutLogOutput)
	APIGetUserWorkoutLogs(input *model.APIGetUserWorkoutLogsInput) (output model.APIGetUserWorkoutLogsOutput)
}
