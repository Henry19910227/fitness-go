package workout_set

import (
	workoutSet "github.com/Henry19910227/fitness-go/internal/v2/resolver/workout_set"
	"gorm.io/gorm"
)

func NewController(db *gorm.DB) Controller {
	resolver := workoutSet.NewResolver(db)
	return New(resolver)
}
