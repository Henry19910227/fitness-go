package plan

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/tool"
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/plan"
)

func NewController(gormTool tool.Gorm) Controller {
	resolver := plan.NewResolver(gormTool)
	return New(resolver)
}
