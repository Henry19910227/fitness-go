package banner


type IDField struct {
	ID *int64 `json:"id,omitempty" gorm:"column:id" example:"1"` //id
}
type CourseIDField struct {
	CourseID *int64 `json:"course_id,omitempty" gorm:"column:course_id" example:"10"` //課表id
}
type UserIDField struct {
	UserID *int64 `json:"user_id,omitempty" gorm:"column:user_id" example:"10001"` //用戶id
}
type ImageField struct {
	Image *string `json:"image,omitempty" gorm:"column:image" example:"1234.jpg"` //圖片
}
type TypeField struct {
	Type *int `json:"type,omitempty" gorm:"column:type;default:1" example:"1"` //類型(1:課表/2:教練/3:訂閱)
}
type CreateAtField struct {
	CreateAt *string `json:"create_at,omitempty" gorm:"column:create_at" example:"2022-06-14 00:00:00"` //創建時間
}
type UpdateAtField struct {
	UpdateAt *string `json:"update_at,omitempty" gorm:"column:update_at" example:"2022-06-14 00:00:00"` //更新時間
}

type Table struct {
	IDField
	CourseIDField
	UserIDField
	ImageField
	TypeField
	CreateAtField
	UpdateAtField
}
func (Table) TableName() string {
	return "banners"
}