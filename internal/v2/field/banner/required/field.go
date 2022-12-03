package required

type IDField struct {
	ID int64 `json:"id" form:"id" binding:"required" example:"1"` //id
}
type CourseIDField struct {
	CourseID int64 `json:"course_id" form:"course_id" binding:"required" example:"10"` //課表id
}
type UserIDField struct {
	UserID int64 `json:"user_id" form:"user_id" binding:"required" example:"10001"` //用戶id
}
type ImageField struct {
	Image string `json:"image" form:"image" binding:"required" example:"1234.jpg"` //圖片
}
type TypeField struct {
	Type int `json:"type" form:"type" binding:"required,oneof=1 2 3" example:"1"` //類型(1:課表/2:教練/3:訂閱)
}
type CreateAtField struct {
	CreateAt string `json:"create_at" gorm:"column:create_at" binding:"required" example:"2022-06-14 00:00:00"` //創建時間
}
type UpdateAtField struct {
	UpdateAt string `json:"update_at" gorm:"column:update_at" binding:"required" example:"2022-06-14 00:00:00"` //更新時間
}