package required

type IDField struct {
	ID int64 `json:"id" gorm:"column:id" binding:"required" example:"1"` //相片id
}
type UserIDField struct {
	UserID int64 `json:"user_id" gorm:"column:user_id" binding:"required" example:"10001"` // 用戶id
}
type PhotoField struct {
	Photo string `json:"photo" gorm:"column:photo" binding:"required" example:"123.jpg"` // 照片
}
type CreateAtField struct {
	CreateAt string `json:"create_at" gorm:"column:create_at" binding:"required" example:"2022-06-14 00:00:00"` //創建時間
}
