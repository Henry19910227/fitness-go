package plan

import (
	"github.com/Henry19910227/fitness-go/internal/v2/service/course"
	"github.com/Henry19910227/fitness-go/internal/v2/service/plan"
	"gorm.io/gorm"
)

func NewResolver(db *gorm.DB) Resolver {
	planService := plan.NewService(db)
	courseService := course.NewService(db)
	return New(planService, courseService)
}
