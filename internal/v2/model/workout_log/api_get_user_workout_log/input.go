package api_get_user_workout_log

import (
	userRequired "github.com/Henry19910227/fitness-go/internal/v2/field/user/required"
	"github.com/Henry19910227/fitness-go/internal/v2/field/workout_log/required"
)

// Input /v2/user/workout_log/{workout_log_id} [GET]
type Input struct {
	userRequired.UserIDField
	Uri Uri
}
type Uri struct {
	required.WorkoutLogIDField
}
