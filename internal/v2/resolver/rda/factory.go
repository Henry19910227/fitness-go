package rda

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/bmr"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/food_calorie"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/tdee"
	"github.com/Henry19910227/fitness-go/internal/v2/service/diet"
	"github.com/Henry19910227/fitness-go/internal/v2/service/rda"
	"gorm.io/gorm"
)

func NewResolver(db *gorm.DB) Resolver {
	rdaService := rda.NewService(db)
	dietService := diet.NewService(db)
	tdeeTool := tdee.NewTool()
	bmrTool := bmr.NewTool()
	calorieTool := food_calorie.NewTool()
	return New(rdaService, dietService, tdeeTool, bmrTool, calorieTool)
}
