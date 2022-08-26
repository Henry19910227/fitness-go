package workout

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/workout"
	"gorm.io/gorm"
)

type Resolver interface {
	APICreateUserWorkout(tx *gorm.DB, input *model.APICreateUserWorkoutInput) (output model.APICreateUserWorkoutOutput)
	APIDeleteUserWorkout(tx *gorm.DB, input *model.APIDeleteUserWorkoutInput) (output model.APIDeleteUserWorkoutOutput)
}
