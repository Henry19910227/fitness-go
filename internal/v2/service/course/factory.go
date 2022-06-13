package course

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/tool"
	"github.com/Henry19910227/fitness-go/internal/v2/repository/course"
)

func NewService(gormTool tool.Gorm) Service {
	courseRepo := course.New(gormTool)
	return New(courseRepo)
}
