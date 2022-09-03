package workout_set_order

import (
	"github.com/Henry19910227/fitness-go/internal/v2/repository/workout_set_order"
	"gorm.io/gorm"
)

func NewService(db *gorm.DB) Service {
	repository := workout_set_order.New(db)
	return New(repository)
}
