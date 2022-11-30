package required

type IDField struct {
	ID int64 `json:"id" gorm:"column:id" binding:"required" example:"2"` // 紀錄 id
}
type UserIDField struct {
	UserID int64 `json:"user_id" gorm:"column:user_id" binding:"required" example:"10001"` // 用戶id
}
type TrainerStatusField struct {
	TrainerStatus int `json:"trainer_status" gorm:"column:trainer_status" binding:"required" example:"1"` // 教練帳戶狀態 (1:正常/2:審核中/3:停權)
}
type CommentField struct {
	Comment string `json:"comment" gorm:"column:comment" binding:"required,max=500" example:"教練看起來不專業"` // 註解
}
type CreateAtField struct {
	CreateAt string `json:"create_at" gorm:"column:create_at" binding:"required" example:"2022-06-12 00:00:00"` // 創建時間
}
