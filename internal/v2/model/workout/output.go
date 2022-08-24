package workout

import "github.com/Henry19910227/fitness-go/internal/v2/model/base"

type Output struct {
	Table
}

func (Output) TableName() string {
	return "workouts"
}

// APICreatePersonalWorkoutOutput /v2/personal/plan/{plan_id}/workout [POST]
type APICreatePersonalWorkoutOutput struct {
	base.Output
	Data *APICreatePersonalWorkoutData `json:"data,omitempty"`
}
type APICreatePersonalWorkoutData struct {
	IDField
}
