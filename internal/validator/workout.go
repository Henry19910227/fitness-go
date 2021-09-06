package validator

type WorkoutIDUri struct {
	WorkoutID int64 `uri:"workout_id" binding:"required" example:"1"`
}

type CreateWorkoutBody struct {
	Name string `json:"name" binding:"required,min=1,max=20" example:"胸肌訓練"`
	WorkoutTemplateID *int64 `json:"workout_template_id" binding:"omitempty" example:"1"` // 訓練模板ID
}

type UpdateWorkoutBody struct {
	Name *string `json:"name" binding:"omitempty,min=1,max=20" example:"胸肌訓練"`
	Equipment *string `json:"equipment" binding:"omitempty,equipment,min=0,max=10" example:"2,3,7"` // 所需器材(1:無需任何器材/2:啞鈴/3:槓鈴/4:固定式器材/5:彈力繩/6:壺鈴/7:訓練椅/8:瑜珈墊/9:其他)
}
