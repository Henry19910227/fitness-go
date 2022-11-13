package subscribe_plan

import (
	"github.com/Henry19910227/fitness-go/internal/v2/service/subscribe_plan"
	"gorm.io/gorm"
)

func NewResolver(db *gorm.DB) Resolver {
	subscribePlanService := subscribe_plan.NewService(db)
	return New(subscribePlanService)
}
