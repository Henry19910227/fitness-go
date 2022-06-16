package food

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/logger"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/orm"
	foodService "github.com/Henry19910227/fitness-go/internal/v2/service/food"
	"github.com/spf13/viper"
)

func NewResolver(gormTool orm.Tool, vp *viper.Viper) Resolver {
	foodSvc := foodService.NewService(gormTool)
	logTool := logger.NewTool(vp)
	return New(foodSvc, logTool)
}
