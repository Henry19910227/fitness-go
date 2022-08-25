package workout

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/workout"
	"gorm.io/gorm"
)

type Resolver interface {
	APICreatePersonalWorkout(tx *gorm.DB, input *model.APICreatePersonalWorkoutInput) (output model.APICreatePersonalWorkoutOutput)
	APIDeletePersonalWorkout(tx *gorm.DB, input *model.APIDeletePersonalWorkoutInput) (output model.APIDeletePersonalWorkoutOutput)
}
