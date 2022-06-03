package validator

type CreateFoodBody struct {
	FoodCategoryID  int64 `json:"food_category_id" binding:"required" example:"1"` //食物類別id
	Name string `json:"name" binding:"required,max=50" example:"地瓜"` //食物名稱
	AmountDesc string `json:"amount_desc" binding:"required" example:"一份地瓜100克70卡"` //份量描述
}
