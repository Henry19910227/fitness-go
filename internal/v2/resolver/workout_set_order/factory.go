package workout_set_order

import (
	"github.com/Henry19910227/fitness-go/internal/v2/service/course"
	"github.com/Henry19910227/fitness-go/internal/v2/service/workout_set_order"
	"gorm.io/gorm"
)

func NewResolver(db *gorm.DB) Resolver {
	workoutSetOrderService := workout_set_order.NewService(db)
	courseService := course.NewService(db)
	return New(workoutSetOrderService, courseService)
}
