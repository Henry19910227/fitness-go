package optional

type IDField struct {
	ID *int64 `json:"id,omitempty" gorm:"column:id" binding:"omitempty" example:"2"` // 紀錄 id
}
type UserIDField struct {
	UserID *int64 `json:"user_id,omitempty" gorm:"column:user_id" binding:"omitempty" example:"10001"` // 用戶id
}
type TrainerStatusField struct {
	TrainerStatus *int `json:"trainer_status,omitempty" gorm:"column:trainer_status" binding:"omitempty" example:"1"` // 教練帳戶狀態 (1:正常/2:審核中/3:停權)
}
type CommentField struct {
	Comment *string `json:"comment,omitempty" gorm:"column:comment" binding:"omitempty,max=500" example:"教練看起來不專業"` // 註解
}
type CreateAtField struct {
	CreateAt *string `json:"create_at,omitempty" gorm:"column:create_at" binding:"omitempty" example:"2022-06-12 00:00:00"` // 創建時間
}
