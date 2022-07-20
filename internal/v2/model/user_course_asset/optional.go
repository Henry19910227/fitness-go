package user_course_asset

type IDOptional struct {
	ID *int64 `json:"id,omitempty" example:"1"` //id
}
type UserIDOptional struct {
	UserID *int64 `json:"user_id,omitempty" example:"10001"` // 用戶id
}
type CourseIDOptional struct {
	CourseID *int64 `json:"course_id,omitempty" example:"10"` //課表id
}
type AvailableOptional struct {
	Available *int `json:"available,omitempty" example:"1"` // 是否可用(0:不可用/1:可用)
}
type CreateAtOptional struct {
	CreateAt *string `json:"create_at,omitempty" example:"2022-06-12 00:00:00"` // 創建時間
}
type UpdateAtOptional struct {
	UpdateAt *string `json:"update_at,omitempty" example:"2022-06-12 00:00:00"` // 更新時間
}
