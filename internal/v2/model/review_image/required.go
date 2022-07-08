package review_image

type IDRequired struct {
	ID int64 `json:"id" uri:"review_image_id" binding:"required" example:"1"` //圖片id
}
type ReviewIDRequired struct {
	ID int64 `json:"review_id" binding:"required" example:"1"` //評論id
}
type ImageRequired struct {
	Image string `json:"image" binding:"required" example:"1234.jpg"` //圖片
}
