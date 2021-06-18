package validator

type CreateWorkoutBody struct {
	Name string `json:"name" binding:"required,min=1,max=20" example:"胸肌訓練"`
}
