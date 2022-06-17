package food

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/tool"
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/food"
)

func NewController(gormTool tool.Gorm) Controller {
	resolver := food.NewResolver(gormTool)
	return New(resolver)
}
