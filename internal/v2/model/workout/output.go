package workout

import "github.com/Henry19910227/fitness-go/internal/v2/model/base"

type Output struct {
	Table
}

func (Output) TableName() string {
	return "workouts"
}

// APICreateUserWorkoutOutput /v2/user/plan/{plan_id}/workout [POST]
type APICreateUserWorkoutOutput struct {
	base.Output
	Data *APICreateUserWorkoutData `json:"data,omitempty"`
}
type APICreateUserWorkoutData struct {
	IDField
}

// APIDeleteUserWorkoutOutput /v2/user/workout/{workout_id} [DELETE]
type APIDeleteUserWorkoutOutput struct {
	base.Output
}
