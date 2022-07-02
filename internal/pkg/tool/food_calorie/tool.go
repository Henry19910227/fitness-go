package food_calorie

import (
	foodCategory "github.com/Henry19910227/fitness-go/internal/v2/model/food_category"
)

type tool struct {
	foodCalorieMap map[int]int
}

func New() Tool {
	foodCalorieMap := map[int]int{
		foodCategory.Grain:     70,
		foodCategory.Meat:      60,
		foodCategory.Fruit:     60,
		foodCategory.Vegetable: 25,
		foodCategory.Dairy:     120,
		foodCategory.Nut:       45,
	}
	return &tool{foodCalorieMap: foodCalorieMap}
}

func (t *tool) Calorie(tag int) int {
	if _, ok := t.foodCalorieMap[tag]; !ok {
		return 0
	}
	return t.foodCalorieMap[tag]
}
