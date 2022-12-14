package food_calorie

import (
	foodCategory "github.com/Henry19910227/fitness-go/internal/v2/model/food_category"
)

const (
	DietTargetLoseFat     = 1 // 減脂
	DietTargetBuildMuscle = 2 // 增肌
	DietTargetKeep        = 3 // 維持健康生活
	DietTargetPowerUp     = 4 // 提升體能與力量
	DietTargetFeed        = 5 // 哺乳者
	DietTargetPregnant    = 6 // 懷孕者
)

const (
	DietTypeGeneral       = 1 // 標準飲食
	DietTypeAllVegan      = 2 // 全素食
	DietTypeOvoLactoVegan = 3 // 蛋奶素食
	DietTypeOvoVegan      = 4 // 蛋素食
	DietTypeLactoVegan    = 5 // 奶素食
)

type tool struct {
	dietTargetMap  map[int]float64
	foodCalorieMap map[int]int
}

func New() Tool {
	dietTargetMap := map[int]float64{
		DietTargetLoseFat:     0.85,
		DietTargetBuildMuscle: 1.15,
		DietTargetKeep:        1,
		DietTargetPowerUp:     1.15,
	}
	foodCalorieMap := map[int]int{
		foodCategory.Grain:     70,
		foodCategory.Meat:      60,
		foodCategory.Fruit:     60,
		foodCategory.Vegetable: 25,
		foodCategory.Dairy:     120,
		foodCategory.Nut:       45,
	}
	return &tool{foodCalorieMap: foodCalorieMap, dietTargetMap: dietTargetMap}
}

func (t *tool) FoodCalorie(tag int) int {
	if _, ok := t.foodCalorieMap[tag]; !ok {
		return 0
	}
	return t.foodCalorieMap[tag]
}

func (t *tool) TargetCalorie(tdee int, target int) float64 {
	if target == DietTargetFeed {
		return float64(tdee) + 600
	}
	if target == DietTargetPregnant {
		return float64(tdee) + 300
	}
	return float64(tdee) * t.dietTargetMap[target]
}
