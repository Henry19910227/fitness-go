package required

type UserIDField struct {
	UserID int64 `json:"user_id" gorm:"column:user_id" binding:"required" example:"10001"` // 用戶 id
}
type CourseIDField struct {
	CourseID int64 `json:"course_id" uri:"course_id" gorm:"column:course_id" binding:"required" example:"10"` //課表id
}
type CreateAtField struct {
	CreateAt string `json:"create_at" gorm:"column:create_at" binding:"required" example:"2022-06-12 00:00:00"` // 創建時間
}