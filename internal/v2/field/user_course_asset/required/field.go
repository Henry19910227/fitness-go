package required

type IDField struct {
	ID int64 `json:"id" gorm:"column:id" binding:"required" example:"1"` //id
}
type UserIDField struct {
	UserID int64 `json:"user_id" gorm:"column:user_id" binding:"required" example:"10001"` // 用戶id
}
type CourseIDField struct {
	CourseID int64 `json:"course_id" gorm:"column:course_id" binding:"required" example:"10"` //課表id
}
type AvailableField struct {
	Available int `json:"available" gorm:"column:available" binding:"required" example:"1"` // 是否可用(0:不可用/1:可用)
}
type SourceField struct {
	Source int `json:"source" gorm:"column:source" binding:"required" example:"1"` // 來源(0:未知/1:購買/2:贈送)
}
type CreateAtField struct {
	CreateAt string `json:"create_at" gorm:"column:create_at" binding:"required" example:"2022-06-12 00:00:00"` // 創建時間
}
type UpdateAtField struct {
	UpdateAt string `json:"update_at" gorm:"column:update_at" binding:"required" example:"2022-06-12 00:00:00"` // 更新時間
}
