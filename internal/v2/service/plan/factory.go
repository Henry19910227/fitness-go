package plan

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/tool"
	"github.com/Henry19910227/fitness-go/internal/v2/repository/plan"
)

func NewService(gormTool tool.Gorm) Service {
	repository := plan.New(gormTool)
	return New(repository)
}
