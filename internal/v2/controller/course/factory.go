package course

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/tool"
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/course"
)

func NewController(gormTool tool.Gorm) Controller {
	resolver := course.NewResolver(gormTool)
	return New(resolver)
}
