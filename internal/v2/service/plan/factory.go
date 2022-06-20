package plan

import (
	"github.com/Henry19910227/fitness-go/internal/v2/repository/plan"
	"gorm.io/gorm"
)

func NewService(db *gorm.DB) Service {
	repository := plan.New(db)
	return New(repository)
}
