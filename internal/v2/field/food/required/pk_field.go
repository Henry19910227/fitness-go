package required

type FoodIDField struct {
	FoodID int64 `json:"food_id" gorm:"column:food_id" binding:"required" example:"1"` //主鍵id
}
