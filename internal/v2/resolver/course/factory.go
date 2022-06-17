package course

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/logger"
	courseService "github.com/Henry19910227/fitness-go/internal/v2/service/course"
	"gorm.io/gorm"
)

func NewResolver(db *gorm.DB) Resolver {
	courseSvc := courseService.NewService(db)
	logTool := logger.NewTool()
	return New(courseSvc, logTool)
}
