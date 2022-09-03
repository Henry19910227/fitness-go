package workout_set_order

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/workout_set_order"
	"gorm.io/gorm"
)

type Resolver interface {
	APIUpdateUserWorkoutSetOrders(tx *gorm.DB, input *model.APIUpdateUserWorkoutSetOrdersInput) (output model.APIUpdateUserWorkoutSetOrdersOutput)
}
