package plan

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/tool"
	planService "github.com/Henry19910227/fitness-go/internal/v2/service/plan"
)

func NewResolver(gormTool tool.Gorm) Resolver {
	planSvc := planService.NewService(gormTool)
	return New(planSvc)
}
