package dto

type Meal struct {
	ID     int64   `json:"id" example:"1"`       //餐食id
	Type   int     `json:"type" example:"3"`     //類型(1:/早餐/2:午餐/3:晚餐/4:點心)
	Amount float64 `json:"amount" example:"0.5"` //數量
	Food   *Food   `json:"food,omitempty"`       //食物
}

type MealType struct {
	Type int `json:"type" example:"3"` //類型(1:/早餐/2:午餐/3:晚餐/4:點心)
}

type MealParamItem struct {
	DietID int64 `json:"diet_id" gorm:"column:diet_id"` //飲食紀錄id
	FoodID int64 `json:"food_id" gorm:"column:food_id"` //食物id
	Type   int   `json:"type" gorm:"column:type"`       //類型(1:/早餐/2:午餐/3:晚餐/4:點心)
	Amount int   `json:"amount" gorm:"column:amount"`   //數量
}

type CreateMealsParam struct {
	MealParamItems []*MealParamItem
}
