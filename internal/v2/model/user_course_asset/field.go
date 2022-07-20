package user_course_asset


type IDField struct {
	ID *int64 `json:"id,omitempty" gorm:"column:id" example:"1"` //id
}
type UserIDField struct {
	UserID *int64 `json:"user_id,omitempty" gorm:"column:user_id" example:"10001"` // 用戶id
}
type CourseIDField struct {
	CourseID *int64 `json:"course_id,omitempty" gorm:"column:course_id" example:"10"` //課表id
}
type AvailableField struct {
	Available *int `json:"available,omitempty" gorm:"column:available" example:"1"` // 是否可用(0:不可用/1:可用)
}
type CreateAtField struct {
	CreateAt *string `json:"create_at,omitempty" gorm:"column:create_at" example:"2022-06-12 00:00:00"` // 創建時間
}
type UpdateAtField struct {
	UpdateAt *string `json:"update_at,omitempty" gorm:"column:update_at" example:"2022-06-12 00:00:00"` // 更新時間
}


type Table struct {
	IDField
	UserIDField
	CourseIDField
	AvailableField
	CreateAtField
	UpdateAtField
}
func (Table) TableName() string {
	return "user_course_assets"
}
