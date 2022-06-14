package service

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/errcode"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool"
	"github.com/Henry19910227/fitness-go/internal/v1/repository"
)

func NewMockRdaService(gormTool tool.Gorm) RDA {
	return NewRDA(repository.NewRDA(gormTool), repository.NewDiet(gormTool), repository.NewTransaction(gormTool),
		tool.NewBMR(), tool.NewTDEE(), tool.NewCalorie(), errcode.NewMockHandler())
}

func NewMockDietService(gormTool tool.Gorm) Diet {
	return NewDiet(repository.NewDiet(gormTool), repository.NewRDA(gormTool), errcode.NewMockHandler())
}

func NewMockMealService(gormTool tool.Gorm) Meal {
	return NewMeal(repository.NewMeal(gormTool), errcode.NewMockHandler())
}
