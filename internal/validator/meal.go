package validator

type MealParamItem struct {
	DietID int64 `json:"diet_id" binding:"required" example:"1"` //飲食紀錄id
	FoodID int64 `json:"food_id" binding:"required" example:"1"` //食物id
	Type int `json:"type" binding:"required,oneof=1 2 3 4" example:"2"` //類型(1:/早餐/2:午餐/3:晚餐/4:點心)
	Amount int `json:"amount" binding:"required" example:"2"` //數量
}
