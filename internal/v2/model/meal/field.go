package meal

type IDField struct {
	ID *int64 `json:"id,omitempty" gorm:"column:id" example:"1"` //餐食id
}
type DietIDField struct {
	DietID *int64 `json:"diet_id,omitempty" uri:"diet_id" gorm:"column:diet_id" example:"1"` //飲食紀錄id
}
type FoodIDField struct {
	FoodID *int64 `json:"food_id,omitempty" gorm:"column:food_id" example:"1"` //食物id
}
type TypeField struct {
	Type *int `json:"type,omitempty" gorm:"column:type;default:1" example:"1"` //類型(1:/早餐/2:午餐/3:晚餐/4:點心)
}
type AmountField struct {
	Amount *float64 `json:"amount,omitempty" gorm:"column:amount;default:0" example:"1"` //數量
}
type CreateAtField struct {
	CreateAt *string `json:"create_at,omitempty" gorm:"column:create_at;default:2022-06-14 00:00:00" example:"2022-06-14 00:00:00"` //創建時間
}

type Table struct {
	IDField
	DietIDField
	FoodIDField
	TypeField
	AmountField
	CreateAtField
}

func (Table) TableName() string {
	return "meals"
}
