package review_image

type IDOptional struct {
	ID *int64 `json:"id,omitempty" binding:"omitempty" example:"1"` //圖片id
}
type ReviewIDOptional struct {
	ID *int64 `json:"review_id,omitempty" binding:"omitempty" example:"1"` //評論id
}
type ImageOptional struct {
	Image *string `json:"image,omitempty" binding:"omitempty" example:"1234.jpg"` //圖片
}
type CreateAtOptional struct {
	CreateAt *string `json:"create_at,omitempty" binding:"omitempty" example:"2022-06-14 00:00:00"` //創建時間
}
