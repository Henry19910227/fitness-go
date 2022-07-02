package food

type IDRequired struct {
	ID int64 `json:"id,omitempty" uri:"food_id" binding:"required" example:"1"` //主鍵id
}
type UserIDRequired struct {
	UserID int64 `json:"user_id,omitempty" binding:"required" example:"10001"` //用戶id
}
type FoodCategoryIDRequired struct {
	FoodCategoryID int64 `json:"food_category_id,omitempty" binding:"required" example:"1"` //食物類別id
}
type SourceRequired struct {
	Source int `json:"source,omitempty" binding:"required" example:"1"` //來源(1:系統創建食物/2:用戶創建食物)
}
type NameRequired struct {
	Name string `json:"name,omitempty" form:"name" binding:"required" example:"蕃薯"` //食物名稱
}
type CalorieRequired struct {
	Calorie int `json:"calorie,omitempty" binding:"required" example:"300"` //食物熱量
}
type AmountDescRequired struct {
	AmountDesc string `json:"amount_desc,omitempty" binding:"required" example:"一份三百卡"` //份量描述
}
type StatusRequired struct {
	Status int `json:"status,omitempty" binding:"required" example:"1"` //狀態(0:下架/1:上架)
}
type IsDeletedRequired struct {
	IsDeleted int `json:"is_deleted,omitempty" binding:"required" example:"0"` //是否刪除
}
