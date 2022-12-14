package rda

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/rda/api_calculate_rda"
	"github.com/Henry19910227/fitness-go/internal/v2/model/rda/api_update_rda"
	"gorm.io/gorm"
)

type Resolver interface {
	APIUpdateRDA(tx *gorm.DB, input *api_update_rda.Input) (output api_update_rda.Output)
	APICalculateRDA(input *api_calculate_rda.Input) (output api_calculate_rda.Output)

	CalculateBMR(birthday string, weight float64, height float64, bodyFat *int, sex string) int
	CalculateTDEE(bmr int, activityLevel int, exerciseFeqLevel int) int
	CalculateCalorie(tdee int, dietTarget int) int
	CalculateProteinCalorie(calorie int, dietTarget int) int
	CalculateFatCalorie(calorie int, dietTarget int) int
	CalculateCarbsCalorie(calorie int, dietTarget int) int
	CalculateProteinAmount(proteinCal int) int
	CalculateCarbsAmount(carbsCal int) int
	CalculateFatAmount(fatCal int) int
	CalculateDairyAmount(dietType int) int
	CalculateVegetableAmount() int
	CalculateFruitAmount() int
	CalculateGrainAmount(carbsAmt int, dairyAmt int, vegetableAmt int, fruitAmt int) int
	CalculateMeatAmount(proteinAmt int, dairyAmt int, grainAmt int, vegetableAmt int) int
	CalculateNutAmount(fatAmt int, dairyAmt int, meatAmt int) int
}
