package workout_set

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/workout_set"
	"gorm.io/gorm"
)

type Resolver interface {
	APIGetCMSWorkoutSets(input *model.APIGetCMSWorkoutSetsInput) interface{}

	APICreateUserWorkoutSets(tx *gorm.DB, input *model.APICreateUserWorkoutSetsInput) (output model.APICreateUserWorkoutSetsOutput)
	APICreateUserWorkoutSetByDuplicate(tx *gorm.DB, input *model.APICreateUserWorkoutSetByDuplicateInput) (output model.APICreateUserWorkoutSetByDuplicateOutput)
	APICreateUserRestSet(input *model.APICreateUserRestSetInput) (output model.APICreateUserRestSetOutput)
	APIDeleteUserWorkoutSet(tx *gorm.DB, input *model.APIDeleteUserWorkoutSetInput) (output model.APIDeleteUserWorkoutSetOutput)
	APIUpdateUserWorkoutSet(tx *gorm.DB, input *model.APIUpdateUserWorkoutSetInput) (output model.APIUpdateUserWorkoutSetOutput)
	APIDeleteUserWorkoutSetStartAudio(input *model.APIDeleteUserWorkoutSetStartAudioInput) (output model.APIDeleteUserWorkoutSetStartAudioOutput)
	APIDeleteUserWorkoutSetProgressAudio(input *model.APIDeleteUserWorkoutSetProgressAudioInput) (output model.APIDeleteUserWorkoutSetProgressAudioOutput)
	APIGetUserWorkoutSets(input *model.APIGetUserWorkoutSetsInput) (output model.APIGetUserWorkoutSetsOutput)

	APIGetTrainerWorkoutSets(input *model.APIGetTrainerWorkoutSetsInput) (output model.APIGetTrainerWorkoutSetsOutput)
	APICreateTrainerWorkoutSets(tx *gorm.DB, input *model.APICreateTrainerWorkoutSetsInput) (output model.APICreateTrainerWorkoutSetsOutput)
	APICreateTrainerRestSet(input *model.APICreateTrainerRestSetInput) (output model.APICreateTrainerRestSetOutput)
	APIDeleteTrainerWorkoutSet(tx *gorm.DB, input *model.APIDeleteTrainerWorkoutSetInput) (output model.APIDeleteTrainerWorkoutSetOutput)
	APIUpdateTrainerWorkoutSet(tx *gorm.DB, input *model.APIUpdateTrainerWorkoutSetInput) (output model.APIUpdateTrainerWorkoutSetOutput)
	APIDeleteTrainerWorkoutSetStartAudio(input *model.APIDeleteTrainerWorkoutSetStartAudioInput) (output model.APIDeleteTrainerWorkoutSetStartAudioOutput)
}
