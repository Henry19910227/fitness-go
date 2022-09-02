package workout_set

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/workout_set"
	"gorm.io/gorm"
)

type Resolver interface {
	APICreateUserWorkoutSets(tx *gorm.DB, input *model.APICreateUserWorkoutSetsInput) (output model.APICreateUserWorkoutSetsOutput)
	APIDeleteUserWorkoutSet(tx *gorm.DB, input *model.APIDeleteUserWorkoutSetInput) (output model.APIDeleteUserWorkoutSetOutput)
	APIUpdateUserWorkoutSet(tx *gorm.DB, input *model.APIUpdateUserWorkoutSetInput) (output model.APIUpdateUserWorkoutSetOutput)
	APIGetUserWorkoutSets(input *model.APIGetUserWorkoutSetsInput) (output model.APIGetUserWorkoutSetsOutput)
	APIGetCMSWorkoutSets(input *model.APIGetCMSWorkoutSetsInput) interface{}
}
