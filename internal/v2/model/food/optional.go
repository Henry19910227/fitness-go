package food

type IDOptional struct {
	ID *int64 `json:"id,omitempty" binding:"omitempty" example:"1"` //主鍵id
}
type UserIDOptional struct {
	UserID *int64 `json:"user_id,omitempty" binding:"omitempty" example:"10001"` //用戶id
}
type FoodCategoryIDOptional struct {
	FoodCategoryID *int64 `json:"food_category_id,omitempty" binding:"omitempty" example:"1"` //食物類別id
}
type SourceOptional struct {
	Source *int `json:"source,omitempty" binding:"omitempty" example:"1"` //來源(1:系統創建食物/2:用戶創建食物)
}
type NameOptional struct {
	Name *string `json:"name,omitempty" form:"name" binding:"omitempty" example:"蕃薯"` //食物名稱
}
type CalorieOptional struct {
	Calorie *int `json:"calorie,omitempty" binding:"omitempty" example:"300"` //食物熱量
}
type AmountDescOptional struct {
	AmountDesc *string `json:"amount_desc,omitempty" binding:"omitempty" example:"一份三百卡"` //份量描述
}
type StatusOptional struct {
	Status *int `json:"status,omitempty" binding:"omitempty" example:"1"` //狀態(0:下架/1:上架)
}
type IsDeletedOptional struct {
	IsDeleted *int `json:"is_deleted,omitempty" binding:"omitempty" example:"0"` //是否刪除
}
