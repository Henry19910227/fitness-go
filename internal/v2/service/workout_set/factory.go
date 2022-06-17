package workout_set

import (
	workoutSet "github.com/Henry19910227/fitness-go/internal/v2/repository/workout_set"
	"gorm.io/gorm"
)

func NewService(db *gorm.DB) Service {
	repository := workoutSet.New(db)
	return New(repository)
}
