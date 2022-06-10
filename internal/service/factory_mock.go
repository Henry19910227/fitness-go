package service

import (
	"github.com/Henry19910227/fitness-go/errcode"
	"github.com/Henry19910227/fitness-go/internal/repository"
	"github.com/Henry19910227/fitness-go/internal/tool"
)

func NewMockRdaService(gormTool tool.Gorm) RDA {
	return NewRDA(repository.NewRDA(gormTool), repository.NewDiet(gormTool), repository.NewTransaction(gormTool),
		tool.NewBMR(), tool.NewTDEE(), tool.NewCalorie(), errcode.NewMockHandler())
}

func NewMockDietService(gormTool tool.Gorm) Diet {
	return NewDiet(repository.NewDiet(gormTool), repository.NewRDA(gormTool), errcode.NewMockHandler())
}
