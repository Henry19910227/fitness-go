package required

type IDField struct {
	ID int64 `json:"id" gorm:"column:id" binding:"required" example:"1"` //id
}
type UserIDField struct {
	UserID int64 `json:"user_id" gorm:"column:user_id" binding:"required" example:"10001"` //用戶id
}
type NameField struct {
	Name string `json:"name" gorm:"column:name" binding:"required" example:"A級教練證照"` //證照名稱
}
type ImageField struct {
	Image string `json:"image" gorm:"column:image" binding:"required" example:"1234.jpg"` //圖片
}
type CreateAtField struct {
	CreateAt string `json:"create_at" gorm:"column:create_at" binding:"required" example:"2022-06-14 00:00:00"` //創建時間
}
type UpdateAtField struct {
	UpdateAt string `json:"update_at" gorm:"column:update_at" binding:"required" example:"2022-06-14 00:00:00"` //更新時間
}
