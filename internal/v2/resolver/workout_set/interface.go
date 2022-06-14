package workout_set

import model "github.com/Henry19910227/fitness-go/internal/v2/model/workout_set"

type Resolver interface {
	APIGetCMSWorkoutSets(input *model.APIGetCMSWorkoutSetsInput) interface{}
}
