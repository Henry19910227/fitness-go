package entity

func NewMockSystemFoods() []*Food {
	datas := []map[string]interface{}{
		{"id": 1, "food_category_id": 1, "source": 1, "name": "米飯",
			"calorie": 70, "amount_desc": "米飯一份70卡", "is_deleted": 0,
			"create_at": "2022-01-01 00:00:00", "update_at": "2022-01-01 00:00:00"},
		{"id": 2, "food_category_id": 2, "source": 1, "name": "地瓜",
			"calorie": 70, "amount_desc": "地瓜一份70卡", "is_deleted": 0,
			"create_at": "2022-01-01 00:00:00", "update_at": "2022-01-01 00:00:00"},
	}
	foods := make([]*Food, 0)
	for _, data := range datas {
		food := Food{}
		food.ID = int64(data["id"].(int))
		food.FoodCategoryID = int64(data["food_category_id"].(int))
		food.Source = data["source"].(int)
		food.Name = data["name"].(string)
		food.Calorie = data["calorie"].(int)
		food.AmountDesc = data["amount_desc"].(string)
		food.IsDeleted = data["is_deleted"].(int)
		food.CreateAt = data["create_at"].(string)
		food.UpdateAt = data["update_at"].(string)
		foods = append(foods, &food)
	}
	return foods
}
