package subscribe_plan

import (
	"github.com/Henry19910227/fitness-go/internal/v2/repository/subscribe_plan"
	"gorm.io/gorm"
)

func NewService(db *gorm.DB) Service {
	repository := subscribe_plan.New(db)
	return New(repository)
}
