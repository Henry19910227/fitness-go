package model

type Meal struct {
	ID int64 `json:"id" gorm:"column:id"` //餐食id
	DietID int64 `json:"diet_id" gorm:"column:diet_id"` //飲食紀錄id
	FoodID *int64 `json:"food_id,omitempty" gorm:"column:food_id"` //食物id
	Type int `json:"type" gorm:"column:type"` //類型(1:/早餐/2:午餐/3:晚餐/4:點心)
	Amount int `json:"amount" gorm:"column:amount"` //數量
	CreateAt string `json:"create_at" gorm:"column:create_at"` //創建日期
	Food *Food `json:"food,omitempty" gorm:"foreignkey:id;references:food_id"` //食物
}

func (Meal) TableName() string {
	return "meals"
}

type SaveMealsParam struct {
	MealItems []*MealItem
}

type FindMealsParam struct {
	MealIDs []int64
}

type MealItem struct {
	DietID int64 `json:"diet_id" gorm:"column:diet_id"` //飲食紀錄id
	FoodID int64 `json:"food_id" gorm:"column:food_id"` //食物id
	Amount int `json:"amount" gorm:"column:amount"` //類型(1:/早餐/2:午餐/3:晚餐/4:點心)
}

func (MealItem) TableName() string {
	return "meals"
}

