package workout

import (
	"github.com/Henry19910227/fitness-go/internal/v2/service/course"
	"github.com/Henry19910227/fitness-go/internal/v2/service/plan"
	"github.com/Henry19910227/fitness-go/internal/v2/service/workout"
	"gorm.io/gorm"
)

func NewResolver(db *gorm.DB) Resolver {
	planService := plan.NewService(db)
	courseService := course.NewService(db)
	workoutService := workout.NewService(db)
	return New(workoutService, planService, courseService)
}
