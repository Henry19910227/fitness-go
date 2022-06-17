package food

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/logger"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/orm"
	foodService "github.com/Henry19910227/fitness-go/internal/v2/service/food"
)

func NewResolver(gormTool orm.Tool) Resolver {
	foodSvc := foodService.NewService(gormTool)
	logTool := logger.NewTool()
	return New(foodSvc, logTool)
}
