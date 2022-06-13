package course

import (
	"github.com/Henry19910227/fitness-go/internal/repository/course"
	"github.com/Henry19910227/fitness-go/internal/tool"
)

func NewService(gormTool tool.Gorm) Service {
	courseRepo := course.New(gormTool)
	return New(courseRepo)
}
