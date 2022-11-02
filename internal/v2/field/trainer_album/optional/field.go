package optional

type IDField struct {
	ID *int64 `json:"id,omitempty" gorm:"column:id" binding:"omitempty" example:"1"` //相片id
}
type UserIDField struct {
	UserID *int64 `json:"user_id,omitempty" gorm:"column:user_id" binding:"omitempty" example:"10001"` // 用戶id
}
type PhotoField struct {
	Photo *string `json:"photo,omitempty" gorm:"column:photo" binding:"omitempty" example:"123.jpg"` // 照片
}
type CreateAtField struct {
	CreateAt *string `json:"create_at,omitempty" gorm:"column:create_at" binding:"omitempty" example:"2022-06-14 00:00:00"` //創建時間
}
