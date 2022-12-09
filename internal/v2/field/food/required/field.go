package required

type IDField struct {
	ID int64 `json:"id" gorm:"column:id" binding:"required" example:"1"` //主鍵id
}
type UserIDField struct {
	UserID int64 `json:"user_id" gorm:"column:user_id" binding:"required" example:"10001"` //用戶id
}
type FoodCategoryIDField struct {
	FoodCategoryID int64 `json:"food_category_id" gorm:"column:food_category_id" binding:"required" example:"1"` //食物類別id
}
type SourceField struct {
	Source int `json:"source" gorm:"column:source" binding:"required" example:"1"` //來源(1:系統創建食物/2:用戶創建食物)
}
type NameField struct {
	Name string `json:"name" form:"name" gorm:"column:name" binding:"required,min=1,max=40" example:"蕃薯"` //食物名稱
}
type CalorieField struct {
	Calorie int `json:"calorie" gorm:"column:calorie" binding:"required" example:"300"` //食物熱量
}
type AmountDescField struct {
	AmountDesc string `json:"amount_desc" gorm:"column:amount_desc" binding:"required" example:"一份三百卡"` //份量描述
}
type StatusField struct {
	Status int `json:"status" gorm:"column:status" binding:"required,oneof=0 1" example:"1"` //狀態(0:下架/1:上架)
}
type IsDeletedField struct {
	IsDeleted int `json:"is_deleted" gorm:"column:is_deleted" binding:"required" example:"0"` //是否刪除
}
type CreateAtField struct {
	CreateAt string `json:"create_at" gorm:"column:create_at" binding:"required" example:"2022-06-14 00:00:00"` //創建時間
}
type UpdateAtField struct {
	UpdateAt string `json:"update_at" gorm:"column:update_at" binding:"required" example:"2022-06-14 00:00:00"` //更新時間
}
