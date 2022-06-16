package food

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/orm"
	"github.com/Henry19910227/fitness-go/internal/v2/repository/food"
)

func NewService(gormTool orm.Tool) Service {
	repository := food.New(gormTool)
	return New(repository)
}
