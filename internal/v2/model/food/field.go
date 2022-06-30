package food

type IDField struct {
	ID *int64 `json:"id,omitempty" gorm:"column:id" example:"1"` //主鍵id
}
type UserIDField struct {
	UserID *int64 `json:"user_id,omitempty" gorm:"column:user_id" example:"10001"` //用戶id
}
type FoodCategoryIDField struct {
	FoodCategoryID *int64 `json:"food_category_id,omitempty" gorm:"column:food_category_id" example:"1"` //食物類別id
}
type SourceField struct {
	Source *int `json:"source,omitempty" gorm:"column:source;default:1" example:"1"` //來源(1:系統創建食物/2:用戶創建食物)
}
type NameField struct {
	Name *string `json:"name,omitempty" form:"name" gorm:"column:name;default:''" example:"蕃薯"` //食物名稱
}
type CalorieField struct {
	Calorie *int `json:"calorie,omitempty" gorm:"column:calorie;default:0" example:"300"` //食物熱量
}
type AmountDescField struct {
	AmountDesc *string `json:"amount_desc,omitempty" gorm:"column:amount_desc;default:''" example:"一份三百卡"` //份量描述
}
type StatusField struct {
	Status *int `json:"status,omitempty" gorm:"column:status;default:1" example:"1"` //狀態(0:下架/1:上架)
}
type IsDeletedField struct {
	IsDeleted *int `json:"is_deleted,omitempty" gorm:"column:is_deleted;default:0" example:"0"` //是否刪除
}
type CreateAtField struct {
	CreateAt *string `json:"create_at,omitempty" gorm:"column:create_at" example:"2022-06-14 00:00:00"` //創建時間
}
type UpdateAtField struct {
	UpdateAt *string `json:"update_at,omitempty" gorm:"column:update_at" example:"2022-06-14 00:00:00"` //更新時間
}

type Table struct {
	IDField
	UserIDField
	FoodCategoryIDField
	SourceField
	NameField
	CalorieField
	AmountDescField
	StatusField
	IsDeletedField
	CreateAtField
	UpdateAtField
}

func (Table) TableName() string {
	return "foods"
}
