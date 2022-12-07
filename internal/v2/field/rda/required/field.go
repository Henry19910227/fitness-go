package required

type IDField struct {
	ID int64 `json:"id" gorm:"column:id" binding:"required" example:"2"` // rda id
}
type UserIDField struct {
	UserID int64 `json:"user_id" gorm:"column:user_id" binding:"required" example:"10001"` //用戶id
}
type TDEEField struct {
	TDEE int `json:"tdee" gorm:"column:tdee" binding:"required" example:"2533"` // TDEE
}
type CalorieField struct {
	Calorie int `json:"calorie" gorm:"column:calorie" binding:"required" example:"2913"` // 目標熱量
}
type ProteinField struct {
	Protein int `json:"protein" gorm:"column:protein" binding:"required" example:"146"` // 蛋白質(克)
}
type FatField struct {
	Fat int `json:"fat" gorm:"column:fat" binding:"required" example:"65"` // 脂肪(克)
}
type CarbsField struct {
	Carbs int `json:"carbs" gorm:"column:carbs" binding:"required" example:"437"` // 碳水化合物(克)
}
type GrainField struct {
	Grain int `json:"grain" gorm:"column:grain" binding:"required" example:"25"` // 穀物類(份)
}
type VegetableField struct {
	Vegetable int `json:"vegetable" gorm:"column:vegetable" binding:"required" example:"5"` // 蔬菜類(份)
}
type FruitField struct {
	Fruit int `json:"fruit" gorm:"column:fruit" binding:"required" example:"2"` // 水果類(份)
}
type MeatField struct {
	Meat int `json:"meat" gorm:"column:meat" binding:"required" example:"12"` // 蛋豆魚肉類(份)
}
type DairyField struct {
	Dairy int `json:"dairy" gorm:"column:dairy" binding:"required" example:"1"` // 乳製品類(份)
}
type NutField struct {
	Nut int `json:"nut" gorm:"column:nut" binding:"required" example:"5"` // 堅果類(份)
}
type CreateAtField struct {
	CreateAt string `json:"create_at" gorm:"column:create_at" binding:"required" example:"2022-06-12 00:00:00"` // 創建時間
}
type UpdateAtField struct {
	UpdateAt string `json:"update_at" gorm:"column:update_at" binding:"required" example:"2022-06-12 00:00:00"` // 更新時間
}
