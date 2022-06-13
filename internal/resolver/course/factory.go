package course

import (
	courseService "github.com/Henry19910227/fitness-go/internal/service/course"
	"github.com/Henry19910227/fitness-go/internal/tool"
)

func NewResolver(gormTool tool.Gorm) Resolver {
	courseSvc := courseService.NewService(gormTool)
	return New(courseSvc)
}
