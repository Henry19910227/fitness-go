package optional

type UserIDField struct {
	UserID *int64 `json:"user_id,omitempty" gorm:"column:user_id" binding:"omitempty" example:"10001"` // 用戶 id
}
type TrainerIDField struct {
	TrainerID *int64 `json:"trainer_id,omitempty" uri:"user_id" gorm:"column:trainer_id" binding:"omitempty" example:"10002"` //教練id
}
type CreateAtField struct {
	CreateAt *string `json:"create_at,omitempty" gorm:"column:create_at" binding:"omitempty" example:"2022-06-12 00:00:00"` // 創建時間
}
