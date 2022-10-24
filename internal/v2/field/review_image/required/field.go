package required

type IDField struct {
	ID int64 `json:"id" uri:"review_image_id" gorm:"column:id" binding:"required" example:"1"` //圖片id
}
type ReviewIDField struct {
	ReviewID int64 `json:"review_id" gorm:"column:review_id" binding:"required" example:"1"` //評論id
}
type ImageField struct {
	Image string `json:"image" gorm:"column:image" binding:"required" example:"1234.jpg"` //圖片
}
type CreateAtField struct {
	CreateAt string `json:"create_at" gorm:"column:create_at" binding:"required" example:"2022-06-14 00:00:00"` //創建時間
}
