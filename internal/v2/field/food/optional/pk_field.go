package optional

type FoodIDField struct {
	FoodID *int64 `json:"food_id,omitempty" gorm:"column:food_id" binding:"omitempty" example:"1"` //主鍵id
}
