package optional

type IDField struct {
	ID *int64 `json:"id,omitempty" gorm:"column:id" binding:"omitempty" example:"2"` // rda id
}
type UserIDField struct {
	UserID *int64 `json:"user_id,omitempty" gorm:"column:user_id" binding:"omitempty" example:"10001"` //用戶id
}
type TDEEField struct {
	TDEE *int `json:"tdee,omitempty" gorm:"column:tdee" binding:"omitempty" example:"2533"` // TDEE
}
type CalorieField struct {
	Calorie *int `json:"calorie,omitempty" gorm:"column:calorie" binding:"omitempty" example:"2913"` // 目標熱量
}
type ProteinField struct {
	Protein *int `json:"protein,omitempty" gorm:"column:protein" binding:"omitempty" example:"146"` // 蛋白質(克)
}
type FatField struct {
	Fat *int `json:"fat,omitempty" gorm:"column:fat" binding:"omitempty" example:"65"` // 脂肪(克)
}
type CarbsField struct {
	Carbs *int `json:"carbs,omitempty" gorm:"column:carbs" binding:"omitempty" example:"437"` // 碳水化合物(克)
}
type GrainField struct {
	Grain *int `json:"grain,omitempty" gorm:"column:grain" binding:"omitempty" example:"25"` // 穀物類(份)
}
type VegetableField struct {
	Vegetable *int `json:"vegetable,omitempty" gorm:"column:vegetable" binding:"omitempty" example:"5"` // 蔬菜類(份)
}
type FruitField struct {
	Fruit *int `json:"fruit,omitempty" gorm:"column:fruit" binding:"omitempty" example:"2"` // 水果類(份)
}
type MeatField struct {
	Meat *int `json:"meat,omitempty" gorm:"column:meat" binding:"omitempty" example:"12"` // 蛋豆魚肉類(份)
}
type DairyField struct {
	Dairy *int `json:"dairy,omitempty" gorm:"column:dairy" binding:"omitempty" example:"1"` // 乳製品類(份)
}
type NutField struct {
	Nut *int `json:"nut,omitempty" gorm:"column:nut" binding:"omitempty" example:"5"` // 堅果類(份)
}
type CreateAtField struct {
	CreateAt *string `json:"create_at,omitempty" gorm:"column:create_at" binding:"omitempty" example:"2022-06-12 00:00:00"` // 創建時間
}
type UpdateAtField struct {
	UpdateAt *string `json:"update_at,omitempty" gorm:"column:update_at" binding:"omitempty" example:"2022-06-12 00:00:00"` // 更新時間
}
