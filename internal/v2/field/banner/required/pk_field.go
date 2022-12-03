package required

type BannerIDField struct {
	BannerID int64 `json:"banner_id" uri:"banner_id" form:"banner_id" binding:"required" example:"1"` //id
}
