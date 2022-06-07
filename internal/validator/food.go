package validator

type FoodIDUri struct {
	FoodID int64 `uri:"food_id" binding:"required" example:"1"` //食物id
}

type CreateFoodBody struct {
	FoodCategoryID  int64 `json:"food_category_id" binding:"required" example:"1"` //食物類別id
	Name string `json:"name" binding:"required,max=50" example:"地瓜"` //食物名稱
	AmountDesc string `json:"amount_desc" binding:"required" example:"一份地瓜100克70卡"` //份量描述
}

type GetFoodsQuery struct {
	FoodCategoryTag int `form:"food_category_tag" binding:"required,oneof=1 2 3 4 5 6" example:"1"` //食物六大類Tag(1:全穀雜糧/2:蛋豆魚肉/3:水果/4:蔬菜/5:乳製品/6:油脂堅果)
}
