package course

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/tool"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/logger"
	courseService "github.com/Henry19910227/fitness-go/internal/v2/service/course"
	"github.com/spf13/viper"
)

func NewResolver(gormTool tool.Gorm, vp *viper.Viper) Resolver {
	courseSvc := courseService.NewService(gormTool)
	logTool := logger.NewTool(vp)
	return New(courseSvc, logTool)
}
