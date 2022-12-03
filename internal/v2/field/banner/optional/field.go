package optional

type IDField struct {
	ID *int64 `json:"id,omitempty" form:"id" binding:"omitempty" example:"1"` //id
}
type CourseIDField struct {
	CourseID *int64 `json:"course_id,omitempty" form:"course_id" binding:"omitempty" example:"10"` //課表id
}
type UserIDField struct {
	UserID *int64 `json:"user_id,omitempty" form:"user_id" binding:"omitempty" example:"10001"` //用戶id
}
type ImageField struct {
	Image *string `json:"image,omitempty" form:"image" binding:"omitempty" example:"1234.jpg"` //圖片
}
type TypeField struct {
	Type *int `json:"type,omitempty" form:"type" binding:"omitempty,oneof=1 2 3" example:"1"` //類型(1:課表/2:教練/3:訂閱)
}
type CreateAtField struct {
	CreateAt *string `json:"create_at,omitempty" gorm:"column:create_at" binding:"omitempty" example:"2022-06-14 00:00:00"` //創建時間
}
type UpdateAtField struct {
	UpdateAt *string `json:"update_at,omitempty" gorm:"column:update_at" binding:"omitempty" example:"2022-06-14 00:00:00"` //更新時間
}
