package course

import (
	"github.com/Henry19910227/fitness-go/internal/resolver/course"
	"github.com/Henry19910227/fitness-go/internal/tool"
)

func NewController(gormTool tool.Gorm) Controller {
	resolver := course.NewResolver(gormTool)
	return New(resolver)
}
