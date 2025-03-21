package workout

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/workout"
	"gorm.io/gorm"
)

type Resolver interface {
	APICreateUserWorkout(tx *gorm.DB, input *model.APICreateUserWorkoutInput) (output model.APICreateUserWorkoutOutput)
	APICreateUserWorkoutFromTemplate(tx *gorm.DB, input *model.APICreateUserWorkoutInput) (output model.APICreateUserWorkoutOutput)
	APIDeleteUserWorkout(tx *gorm.DB, input *model.APIDeleteUserWorkoutInput) (output model.APIDeleteUserWorkoutOutput)
	APIGetUserWorkouts(input *model.APIGetUserWorkoutsInput) (output model.APIGetUserWorkoutsOutput)
	APIUpdateUserWorkout(tx *gorm.DB, input *model.APIUpdateUserWorkoutInput) (output model.APIUpdateUserWorkoutOutput)
	APIDeleteUserWorkoutStartAudio(input *model.APIDeleteUserWorkoutStartAudioInput) (output model.APIDeleteUserWorkoutStartAudioOutput)
	APIDeleteUserWorkoutEndAudio(input *model.APIDeleteUserWorkoutEndAudioInput) (output model.APIDeleteUserWorkoutEndAudioOutput)

	APICreateTrainerWorkout(tx *gorm.DB, input *model.APICreateTrainerWorkoutInput) (output model.APICreateTrainerWorkoutOutput)
	APICreateTrainerWorkoutFromTemplate(tx *gorm.DB, input *model.APICreateTrainerWorkoutInput) (output model.APICreateTrainerWorkoutOutput)
	APIDeleteTrainerWorkout(tx *gorm.DB, input *model.APIDeleteTrainerWorkoutInput) (output model.APIDeleteTrainerWorkoutOutput)
	APIGetTrainerWorkouts(input *model.APIGetTrainerWorkoutsInput) (output model.APIGetTrainerWorkoutsOutput)
	APIUpdateTrainerWorkout(tx *gorm.DB, input *model.APIUpdateTrainerWorkoutInput) (output model.APIUpdateTrainerWorkoutOutput)
	APIDeleteTrainerWorkoutStartAudio(input *model.APIDeleteTrainerWorkoutStartAudioInput) (output model.APIDeleteTrainerWorkoutStartAudioOutput)
	APIDeleteTrainerWorkoutEndAudio(input *model.APIDeleteTrainerWorkoutEndAudioInput) (output model.APIDeleteTrainerWorkoutEndAudioOutput)

	APIGetStoreWorkouts(input *model.APIGetStoreWorkoutsInput) (output model.APIGetStoreWorkoutsOutput)
}
