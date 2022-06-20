package plan

import (
	planService "github.com/Henry19910227/fitness-go/internal/v2/service/plan"
	"gorm.io/gorm"
)

func NewResolver(db *gorm.DB) Resolver {
	planSvc := planService.NewService(db)
	return New(planSvc)
}
