package required

type IDField struct {
	ID int64 `json:"id" gorm:"column:id" binding:"required" example:"1"` //餐食id
}
type DietIDField struct {
	DietID int64 `json:"diet_id" uri:"diet_id" gorm:"column:diet_id" binding:"required" example:"1"` //飲食紀錄id
}
type FoodIDField struct {
	FoodID int64 `json:"food_id" gorm:"column:food_id" binding:"required" example:"1"` //食物id
}
type TypeField struct {
	Type int `json:"type" gorm:"column:type;default:1" binding:"required" example:"1"` //類型(1:/早餐/2:午餐/3:晚餐/4:點心)
}
type AmountField struct {
	Amount float64 `json:"amount" gorm:"column:amount;default:0" binding:"required" example:"1"` //數量
}
type CreateAtField struct {
	CreateAt string `json:"create_at" gorm:"column:create_at;default:2022-06-14 00:00:00" binding:"required" example:"2022-06-14 00:00:00"` //創建時間
}
