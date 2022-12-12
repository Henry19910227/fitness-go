package required

type IDField struct {
	ID int64 `json:"id" uri:"diet_id" gorm:"column:id" binding:"required" example:"1"` //id
}
type UserIDField struct {
	UserID int64 `json:"user_id" gorm:"column:user_id" binding:"required" example:"10001"` //用戶id
}
type RdaIDField struct {
	RdaID int64 `json:"rda_id" gorm:"column:rda_id" binding:"required" example:"1"` //建議營養id
}
type ScheduleAtField struct {
	ScheduleAt string `json:"schedule_at" form:"schedule_at" gorm:"column:schedule_at;default:2022-06-14 00:00:00" binding:"required,datetime=2006-01-02" example:"2022-06-14"` //排程時間
}
type CreateAtField struct {
	CreateAt string `json:"create_at" gorm:"column:create_at;default:2022-06-14 00:00:00" binding:"required" example:"2022-06-14 00:00:00"` //創建時間
}
type UpdateAtField struct {
	UpdateAt string `json:"update_at" gorm:"column:update_at;default:2022-06-14 00:00:00" binding:"required" example:"2022-06-14 00:00:00"` //更新時間
}
