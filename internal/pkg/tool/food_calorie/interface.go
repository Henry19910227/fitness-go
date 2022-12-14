package food_calorie

type Tool interface {
	FoodCalorie(tag int) int
	TargetCalorie(tdee int, target int) float64
}
