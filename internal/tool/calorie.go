package tool

import "github.com/Henry19910227/fitness-go/internal/global"

type calorie struct {
	dietTargetMap map[global.DietTarget]float64
	foodCalorieMap map[global.FoodCategoryTag]int
}

func NewCalorie() Calorie {
	dietTargetMap := map[global.DietTarget]float64{
		global.DietTargetLoseFat:     0.85,
		global.DietTargetBuildMuscle: 1.15,
		global.DietTargetKeep:        1,
		global.DietTargetPowerUp:     1.15,
	}
	foodCalorieMap := map[global.FoodCategoryTag]int{
		global.GrainTag: 70,
		global.MeatTag:  60,
		global.FruitTag: 60,
		global.VegetableTag: 25,
		global.DairyTag: 120,
		global.NutTag: 45,
	}
	return &calorie{dietTargetMap: dietTargetMap, foodCalorieMap: foodCalorieMap}
}

func (c *calorie) TargetCalorie(tdee int, target global.DietTarget) float64 {
	if target == global.DietTargetFeed {
		return float64(tdee) + 600
	}
	if target == global.DietTargetPregnant {
		return float64(tdee) + 300
	}
	return float64(tdee) * c.dietTargetMap[target]
}

func (c *calorie) FoodCalorie(tag global.FoodCategoryTag) int {
	if _, ok := c.foodCalorieMap[tag]; !ok {
		return 0
	}
	return c.foodCalorieMap[tag]
}
