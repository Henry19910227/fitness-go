package course

import (
	"github.com/Henry19910227/fitness-go/internal/tool"
	courseService "github.com/Henry19910227/fitness-go/internal/v2/service/course"
)

func NewResolver(gormTool tool.Gorm) Resolver {
	courseSvc := courseService.NewService(gormTool)
	return New(courseSvc)
}
